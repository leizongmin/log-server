// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import (
	"LogServer/utils"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func Start(options ServerOptions) error {

	var err error

	Options.Dir = options.Dir
	Options.Listen = options.Listen

	if err = setDir(options.Dir); err != nil {
		return err
	}

	if err = listen(options.Listen); err != nil {
		return err
	}

	return nil

}

func listen(addr string) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/log/stream", logStream)

	srv := http.Server{
		Addr:           addr,
		Handler:        mux,
		MaxHeaderBytes: 10240,
	}

	return listenServer(srv)

}

func listenServer(srv http.Server) error {

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
		autoRemoveUnixSocketFile(srv.Addr)
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

func autoRemoveUnixSocketFile(file string) {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT)

	removeFile := func(s os.Signal) {
		log.Printf("get signal %s, trying to remove unix socket file %s before process exit\n", s, file)
		err := os.Remove(file)
		if err != nil {
			log.Printf("fail to remove file: %s\n", err)
		}
		log.Printf("exit\n")
		os.Exit(0)
	}

	go func() {
		s := <-sig
		switch s {
		case os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
			removeFile(s)
		}
	}()

}

func setDir(dir string) error {

	if err := utils.Mkdirp(dir); err != nil {
		return err
	}

	return nil

}
