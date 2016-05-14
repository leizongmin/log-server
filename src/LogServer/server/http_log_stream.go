// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

func logStream(w http.ResponseWriter, r *http.Request) {

	var err error

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	scaner := bufio.NewScanner(r.Body)
	for scaner.Scan() {

		text := scaner.Text()

		log := LogLine{}
		if err = json.Unmarshal([]byte(text), &log); err != nil {
			responseLine(w, fmt.Sprintf("error: %s", err))
			continue
		}

		if log.Path == "" {
			responseLine(w, `error: missing "path"`)
			continue
		}
		if log.ID == "" {
			responseLine(w, `error: missing "id"`)
			continue
		}
		if len(log.Data) < 1 {
			responseLine(w, `error: missing "data"`)
			continue
		}

		if err = WriteLog(log); err != nil {
			responseLine(w, fmt.Sprintf("error: %s", err))
			continue
		}
		responseLine(w, fmt.Sprintf("success: %s", log.ID))

	}

}

func responseLine(w http.ResponseWriter, s string) {

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	fmt.Fprintf(w, s)
	flusher.Flush()

}
