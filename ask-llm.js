#!/usr/bin/env node

const readline = require('readline');

const LLM_API_BASE_URL = process.env.LLM_API_BASE_URL || 'https://api.openai.com/v1';
const LLM_API_KEY = process.env.LLM_API_KEY || process.env.OPENAI_API_KEY;
const LLM_CHAT_MODEL = process.env.LLM_CHAT_MODEL;
const LLM_STREAMING = process.env.LLM_STREAMING !== 'no';

const LLM_DEBUG = process.env.LLM_DEBUG;

/**
 * Represents a chat message.
 *
 * @typedef {Object} Message
 * @property {'system'|'user'|'assistant'} role
 * @property {string} content
 */

/**
 * A callback function to stream then completion.
 *
 * @callback CompletionHandler
 * @param {string} text
 * @returns {void}
 */

/**
 * Generates a chat completion using a RESTful LLM API service.
 *
 * @param {Array<Message>} messages - List of chat messages.
 * @param {CompletionHandler=} handler - An optional callback to stream the completion.
 * @returns {Promise<string>} The completion generated by the LLM.
 */
const chat = async (messages, handler) => {
    const url = `${LLM_API_BASE_URL}/chat/completions`;
    const auth = LLM_API_KEY ? { 'Authorization': `Bearer ${LLM_API_KEY}` } : {};
    const model = LLM_CHAT_MODEL || 'gpt-4.1-nano';
    const stop = ['<|im_end|>', '<|end|>', '<|eot_id|>'];
    const max_tokens = 200;
    const temperature = 0;
    const stream = LLM_STREAMING && typeof handler === 'function';
    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', ...auth },
        body: JSON.stringify({ messages, model, stop, max_tokens, temperature, stream })
    });
    if (!response.ok) {
        throw new Error(`HTTP error with the status: ${response.status} ${response.statusText}`);
    }

    if (!stream) {
        const data = await response.json();
        const { choices } = data;
        const first = choices[0];
        const { message } = first;
        const { content } = message;
        const answer = content.trim();
        handler && handler(answer);
        return answer;
    }

    const parse = (line) => {
        const separator = line.indexOf(':');
        if (separator < 0) {
            return '';
        }
        const key = line.substring(0, separator).trim();
        const payload = line.substring(separator + 1);
        if (key === 'data') {
            let partial = null;
            try {
                const { choices } = JSON.parse(payload);
                const [choice] = choices;
                const { delta } = choice;
                partial = delta?.content;
            } catch (e) {
                // ignore
            } finally {
                return partial;
            }
        } else {
            return '';
        }
    }

    const reader = response.body.getReader();
    const decoder = new TextDecoder();

    let answer = '';
    let buffer = '';
    while (true) {
        const { value, done } = await reader.read();
        if (done) {
            break;
        }
        const lines = decoder.decode(value).split('\n');
        for (let i = 0; i < lines.length; ++i) {
            const line = buffer + lines[i];
            if (line[0] === ':') {
                buffer = '';
                continue;
            }
            if (line === 'data: [DONE]') {
                break;
            }
            if (line.length > 0) {
                const partial = parse(line.trim());
                if (partial === null) {
                    buffer = line;
                } else if (partial && partial.length > 0) {
                    buffer = '';
                    if (answer.length < 1) {
                        const leading = partial.trim();
                        answer = leading;
                        handler && (leading.length > 0) && handler(leading);
                    } else {
                        answer += partial;
                        handler && handler(partial);
                    }
                }
            }
        }
    }
    return answer;
}

const SYSTEM_PROMPT = 'Answer the question politely and concisely.';

(async () => {
    console.log(`Using LLM at ${LLM_API_BASE_URL}.`);
    console.log('Press Ctrl+D to exit.')
    console.log();

    const messages = [];
    messages.push({ role: 'system', content: SYSTEM_PROMPT });

    let loop = true;
    const io = readline.createInterface({ input: process.stdin, output: process.stdout });
    io.on('close', () => { loop = false; });

    const qa = () => {
        io.question('>> ', async (question) => {
            messages.push({ role: 'user', content: question });
            const start = Date.now();
            const answer = await chat(messages, (str) => process.stdout.write(str));
            messages.push({ role: 'assistant', content: answer.trim() });
            console.log();
            const elapsed = Date.now() - start;
            LLM_DEBUG && console.log(`[${elapsed} ms]`);
            console.log();
            loop && qa();
        })
    }

    qa();
})();
