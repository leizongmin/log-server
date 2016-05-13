// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import (
	"bufio"
	"fmt"
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
