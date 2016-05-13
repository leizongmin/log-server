// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import (
	"LogServer/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

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

	file := Options.LogFiles[log.Path]
	if file.Path != log.Path {
		openLogFileForWrite(log.Path, "")
		file = Options.LogFiles[log.Path]
	}

	if file.Path != log.Path {
		return errors.New("failed to open log file")
	}

	text, err := json.Marshal(log.Data)
	if err != nil {
		return err
	}

	_, err = file.Handle.Write(text)
	if err != nil {
		return err
	}

	return nil

}

func updateLogFilesName() {

	for path, file := range Options.LogFiles {

		fileName := getCurrentLogFileName(path)

		if fileName != file.FileName {
			file.Handle.Close()
			openLogFileForWrite(path, fileName)
		}

	}

}

func getCurrentLogFileName(path string) string {

	timeString := utils.GetFormattedTime("Ymd/Ymd-h")
	fileName := fmt.Sprintf("%s/%s/%s.log", Options.Dir, path, timeString)

	return fileName

}

func openLogFileForWrite(path string, fileName string) {

	if fileName == "" {
		fileName = getCurrentLogFileName(path)
	}

	if err := utils.MkdirpByFileName(fileName); err != nil {
		log.Printf("failed to open log file: %s\n", err)
		return
	}

	handle, err := os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Printf("failed to open log file: %s\n", err)
		return
	}

	if oldFile := Options.LogFiles[path]; oldFile.Path != "" {
		oldFile.Handle.Close()
	}

	Options.LogFiles[path] = LogFile{
		Path:     path,
		FileName: fileName,
		Handle:   handle,
	}
	log.Printf("opened new log file: %s\n", fileName)

}

var timer *time.Timer

func init() {

	updateLogFilesName()

	timer = time.NewTimer(time.Second * 10)
	go func() {
		for {
			<-timer.C
			updateLogFilesName()
		}
	}()

}
