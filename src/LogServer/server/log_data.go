// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import "fmt"

// A LogLine describes a log data item
type LogLine struct {
	ID   string                 `json:"id"`
	Path string                 `json:"path"`
	Data map[string]interface{} `json:"data"`
}

// WriteLogFlat writes a log record
func WriteLogFlat(id string, path string, data map[string]interface{}) error {

	return WriteLog(LogLine{
		ID:   id,
		Path: path,
		Data: data,
	})

}

// WriteLog writes a log record with a LogLine struct
func WriteLog(log LogLine) error {

	fmt.Println(log)

	return nil

}
