#!/usr/bin/env node

const readline = require('readline');

const LLM_API_BASE_URL = process.env.LLM_API_BASE_URL || 'https://api.openai.com/v1';
const LLM_API_KEY = process.env.LLM_API_KEY || process.env.OPENAI_API_KEY;
const LLM_CHAT_MODEL = process.env.LLM_CHAT_MODEL;

const LLM_DEBUG = process.env.LLM_DEBUG;

const chat = async (messages) => {
    const url = `${LLM_API_BASE_URL}/chat/completions`;
    const auth = LLM_API_KEY ? { 'Authorization': `Bearer ${LLM_API_KEY}` } : {};
    const model = LLM_CHAT_MODEL || 'gpt-3.5-turbo';
    const max_tokens = 200;
    const temperature = 0;
    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', ...auth },
        body: JSON.stringify({ messages, model, max_tokens, temperature })
    });
    if (!response.ok) {
        throw new Error(`HTTP error with the status: ${response.status} ${response.statusText}`);
    }

    const data = await response.json();
    const { choices } = data;
    const first = choices[0];
    const { message } = first;
    const { content } = message;
    const answer = content.trim();
    return answer;
}

const SYSTEM_PROMPT = 'Answer the question politely and concisely.';

(async () => {
    console.log(`Using LLM at ${LLM_API_BASE_URL}.`);
    console.log('Press Ctrl+D to exit.')
    console.log();

    const messages = [];
    messages.push({ role: 'system', content: SYSTEM_PROMPT });

    const interface = readline.createInterface({ input: process.stdin, output: process.stdout });

    const qa = () => {
        let loop = true;
        interface.on('close', () => { loop = false; });
        interface.question('>> ', async (question) => {
            messages.push({ role: 'user', content: question });
            const start = Date.now();
            const answer = await chat(messages);
            messages.push({ role: 'assistant', content: answer });
            console.log(answer);
            const elapsed = Date.now() - start;
            LLM_DEBUG && console.log(`[${elapsed} ms]`);
            console.log();
            loop && qa();
        })
    }

    qa();
})();