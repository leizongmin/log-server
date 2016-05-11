package main

import (
	"LogServer/server"
	"log"
	"os"
)

func main() {

	log.Printf("PID: %d", os.Getpid())

	addr := "/tmp/logserver.sock"
	//addr := ":8080"
	log.Printf("Listen: trying to listen on %s\n", addr)

	err := server.Listen(addr)
	if err != nil {
		log.Fatal("Listen: ", err)
	}

}
