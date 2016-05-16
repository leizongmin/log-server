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

var optionListen = flag.String("listen", ":8080", `server listen path, e.g ":8080" or "/var/run/logserver.sock"`)
var optionDir = flag.String("dir", "./data", `root directory for logs data`)
var optionsFormat = flag.String("format", "Ymd/Ymd-H", `file name format, e.g "Ymd/Ymd-H"`)
var optionsDuration = flag.Int64("duration", 5, `ticker duration`)

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

	var err error
	var listen string = *optionListen
	var dir string = *optionDir
	var duration int64 = *optionsDuration
	var fileNameFormat = *optionsFormat

	log.Printf("option listen=%s\n", listen)
	log.Printf("option dir=%s\n", dir)
	log.Printf("option duration=%d\n", duration)
	log.Printf("option fileNameFormat=%s\n", fileNameFormat)

	log.Printf("pid: %d", os.Getpid())

	log.Printf("starting server...")
	err = server.Start(server.ServerOptions{
		Listen:         listen,
		Dir:            dir,
		Duration:       duration,
		FileNameFormat: fileNameFormat,
	})
	if err != nil {
		log.Fatal("listen: ", err)
	}

}
