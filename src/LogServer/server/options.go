// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import "os"

type ServerOptions struct {
	Listen   string
	Dir      string
	LogFiles map[string]LogFile
}

type LogFile struct {
	Path     string
	FileName string
	Handle   *os.File
}

// Options stores the config for server
var Options = ServerOptions{}
