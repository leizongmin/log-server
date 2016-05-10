package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"
)

func logStream(w http.ResponseWriter, r *http.Request) {

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	scaner := bufio.NewScanner(r.Body)
	for scaner.Scan() {
		text := scaner.Text()
		fmt.Fprintf(w, "time: %s, data: %s\n", time.UnixDate, text)
		flusher.Flush()
		println(text)
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/log/stream", logStream)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
