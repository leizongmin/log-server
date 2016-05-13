// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package main

import (
	"LogServer/server"
	"flag"
	"fmt"
	"log"
	"os"
)

var optionListen = flag.String("listen", ":8080", `server listen path, e.g. ":8080" or "/var/run/logserver.sock"`)
var optionDir = flag.String("dir", "./data", `root directory for logs data`)

func usage() {

	fmt.Printf("%s\n", `
Usage: logserver [options]

Options:
	`)

	flag.PrintDefaults()
	os.Exit(0)

}

func main() {

	flag.Usage = usage
	flag.Parse()

	log.Printf("option listen=%s\n", *optionListen)
	log.Printf("option dir=%s\n", *optionDir)

	start()

}

func start() {

	var err error
	var addr string = *optionListen
	var dir string = *optionDir

	log.Printf("pid: %d", os.Getpid())

	log.Printf("starting server...")
	err = server.Start(server.ServerOptions{
		Listen: addr,
		Dir:    dir,
	})
	if err != nil {
		log.Fatal("listen: ", err)
	}

}
