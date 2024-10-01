#!/usr/bin/env bb

(require '[clojure.string :as str]
         '[babashka.http-client :as http]
         '[cheshire.core :as json])

(def LLM-API-BASE-URL (or (System/getenv "LLM_API_BASE_URL") "https://api.openai.com/v1"))
(def LLM-API-KEY (or (System/getenv "LLM_API_KEY") (System/getenv "OPENAI_API_KEY")))
(def LLM-CHAT-MODEL (System/getenv "LLM_CHAT_MODEL"))
(def LLM-STREAMING (not= "no" (System/getenv "LLM_STREAMING")))

(def LLM-DEBUG (System/getenv "LLM_DEBUG"))

(defn http-json-headers [bearer]
  (if bearer {:content-type "application/json"
              :authorization (str "Bearer " bearer)}
      {:content-type "application/json"}))

(def LLM-CHAT-URL (str LLM-API-BASE-URL "/chat/completions"))

(def SYSTEM-PROMPT "Answer the question politely and concisely.")

(defonce llm-messages (atom [{:role "system"
                              :content SYSTEM-PROMPT}]))

(defn add-message! [role content]
  (swap! llm-messages conj {:role role :content content}))

(defn make-reader [response]
  (java.io.BufferedReader. (java.io.InputStreamReader. (:body response))))

(defn json-parse [str]
  (json/parse-string str true))

(defn parse-line [line handler]
  (try
    (some-> line json-parse
            :choices first :delta :content
            (#(do (when handler (handler %))
                  (str %))))
    (catch Exception e nil)))

(defn decode-stream [response handler]
  (with-open [reader (make-reader response)]
    (loop [answer ""]
      (if-let [line (.readLine reader)]
        (let [trimmed-line (str/trim line)]
          (cond
            (str/blank? trimmed-line) (recur answer)
            (str/starts-with? trimmed-line "data: ")
            (let [new-answer (parse-line (str/trim (subs trimmed-line 6)) handler)]
              (recur (or new-answer answer)))
            :else (recur answer)))
        answer))))

(defn chat [messages handler]
  (let [stream (and LLM-STREAMING (some? handler))
        payload {:messages messages
                 :model (or LLM-CHAT-MODEL "gpt-4o-mini")
                 :stop ["<|im_end|>" "<|end|>" "<|eot_id|>"]
                 :max_tokens 200
                 :temperature 0
                 :stream stream}
        options {:headers (http-json-headers LLM-API-KEY)
                 :body (json/encode payload)}
        options (if stream (assoc options :as :stream) options)
        response (http/post LLM-CHAT-URL options)]
    (if stream
      (decode-stream response handler)
      (let [body (-> response :body json-parse)
            answer (-> body :choices first :message :content str/trim)]
        (when handler (handler answer))
        answer))))

(defn print-stdout [str]
  (print str)
  (flush))

(defn ask-llm [question]
  (add-message! "user" question)
  (let [answer (chat @llm-messages print-stdout)]
    (add-message! "assistant" answer)
    (println)
    answer))

(defmacro measure-time [f]
  `(let [start-time# (System/currentTimeMillis)
         result# ~f
         end-time# (System/currentTimeMillis)]
     (when LLM-DEBUG (println (str "[" (- end-time# start-time#) " ms]")))
     result#))

(defn qa []
  (loop []
    (print ">> ")
    (flush)
    (let [question (read-line)]
      (when question
        (-> question ask-llm measure-time)
        (println)
        (flush)
        (recur)))))

(defn -main []
  (println "Using LLM at" LLM-API-BASE-URL)
  (println "Press Ctrl+D to exit.")
  (println)
  (qa))

(-main)
