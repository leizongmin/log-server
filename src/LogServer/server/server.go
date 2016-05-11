package server

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
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

func Listen(addr string) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/log/stream", logStream)

	srv := http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 10240,
	}

	return listenAndServe(srv)

}

func listenAndServe(srv http.Server) error {

	var proto string

	if srv.Addr == "" {
		srv.Addr = ":http"
	}

	if strings.Contains(srv.Addr, "/") {
		proto = "unix"
		err := checkUnixSocketFile(srv.Addr)
		if err != nil {
			return err
		}
	} else {
		proto = "tcp"
	}

	l, e := net.Listen(proto, srv.Addr)
	if e != nil {
		return e
	}

	return srv.Serve(l)

}

func checkUnixSocketFile(file string) error {

	exists, err := checkFileExists(file)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(fmt.Sprintf("unix file %s is already exists", file))
	}

	return nil

}

func checkFileExists(file string) (bool, error) {

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}

}
