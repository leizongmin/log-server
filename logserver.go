package main

import (
	"LogServer/server"
	"flag"
	"fmt"
	"log"
	"os"
)

var listenPath = flag.String("listen", ":8080", `server listen path, e.g. ":8080" or "/var/run/logserver.sock"`)

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

	start()

}

func start() {

	log.Printf("PID: %d", os.Getpid())

	addr := *listenPath
	log.Printf("Listen: trying to listen on %s\n", addr)

	err := server.Listen(addr)
	if err != nil {
		log.Fatal("Listen: ", err)
	}

}
