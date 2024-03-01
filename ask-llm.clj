#!/usr/bin/env bb

(require '[clojure.string :as str]
         '[babashka.curl :as curl]
         '[cheshire.core :as json])

(def LLM-API-BASE-URL (or (System/getenv "LLM_API_BASE_URL") "https://api.openai.com/v1"))
(def LLM-API-KEY (or (System/getenv "LLM_API_KEY") (System/getenv "OPENAI_API_KEY")))
(def LLM-CHAT-MODEL (System/getenv "LLM_CHAT_MODEL"))

(def LLM-DEBUG (System/getenv "LLM_DEBUG"))

(defn http-post [url bearer payload]
  (-> (curl/post url {:body (json/generate-string payload {:key-fn (comp name keyword)})
                      :headers {"Content-Type" "application/json"
                                "Authorization" (if bearer (str "Bearer " bearer) nil)}})
      :body (json/parse-string true)))

(def LLM-CHAT-URL (str LLM-API-BASE-URL "/chat/completions"))

(def SYSTEM-PROMPT "Answer the question politely and concisely.")

(defonce llm-messages (atom [{:role "system"
                              :content SYSTEM-PROMPT}]))

(defn add-message! [role content]
  (swap! llm-messages conj {:role role :content content}))

(defn chat [messages]
  (let [body {:messages messages
              :model (or LLM-CHAT-MODEL "gpt-3.5-turbo")
              :max_tokens 200
              :temperature 0}
        response (http-post LLM-CHAT-URL LLM-API-KEY body)]
    (-> response :choices first :message :content str/trim)))

(defn ask-llm [question]
  (add-message! "user" question)
  (let [answer (chat @llm-messages)]
    (add-message! "assistant" answer)
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
        (-> question ask-llm println measure-time)
        (println)
        (flush)
        (recur)))))

(defn -main []
  (println "Using LLM at" LLM-API-BASE-URL)
  (println "Press Ctrl+D to exit.")
  (println)
  (qa))

(-main)
