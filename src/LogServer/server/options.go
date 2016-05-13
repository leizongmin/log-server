// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

type ServerOptions struct {
	Listen string `json:"listen"`
	Dir    string `json:"dir"`
}

// Options stores the config for server
var Options = ServerOptions{}
