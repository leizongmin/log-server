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

	_, err = file.Handle.WriteString("\r\n")
	if err != nil {
		return err
	}

	return nil

}

func syncLogFiles() {

	log.Printf("sync log files...")

	for path, file := range Options.LogFiles {

		fileName := getCurrentLogFileName(path)

		if fileName != file.FileName {
			file.Handle.Sync()
			file.Handle.Close()
			delete(Options.LogFiles, path)
			log.Printf("close log file: %s\n", file.FileName)
		}

	}

}

func getCurrentLogFileName(path string) string {

	timeString := utils.GetFormattedTime("Ymd/Ymd-Hi")
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

	handle, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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

var ticker *time.Ticker

func init() {

	ticker = time.NewTicker(time.Second * 5)
	go func() {
		for {
			<-ticker.C
			syncLogFiles()
		}
	}()

}
