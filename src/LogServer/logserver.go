package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"
)

func logStream(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	scaner := bufio.NewScanner(r.Body)
	for scaner.Scan() {
		text := scaner.Text()
		fmt.Fprintln(w, time.UnixDate, text)
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
