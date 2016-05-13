// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package utils

import (
	"os"
	"strconv"
	"time"
)

// Mkdirp Like `os.MkdirAll()`, but if directory already exists don't returns error
func Mkdirp(dir string) error {

	err := os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil

}

func LeftPadInt(n int, v int) string {

	s := strconv.Itoa(v)
	for i := len(s); i < n; i++ {
		s = "0" + s
	}

	return s

}

// A TimeInfo describes time info
type TimeInfo struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
	Time   *time.Time
}

// GetTime returns a TimeInfo
func GetTime() TimeInfo {

	t := time.Now()
	d := TimeInfo{
		Year:   LeftPadInt(4, t.Year()),
		Month:  LeftPadInt(2, int(t.Month())),
		Day:    LeftPadInt(2, t.Day()),
		Hour:   LeftPadInt(2, t.Hour()),
		Minute: LeftPadInt(2, t.Minute()),
		Second: LeftPadInt(2, t.Second()),
		Time:   &t,
	}

	return d

}

// Format returns a time string with specified format
// e.g `Y-m-d H:i:s`
func (t TimeInfo) Format(f string) string {

	r := []rune(f)
	var s string
	for _, v := range r {
		switch v {
		case 'Y':
			s += t.Year
		case 'm':
			s += t.Month
		case 'd':
			s += t.Day
		case 'H':
			s += t.Hour
		case 'i':
			s += t.Minute
		case 's':
			s += t.Second
		default:
			s += string(v)

		}
	}

	return s

}

// GetTimeFormat returns a time string with specified format
func GetTimeFormat(f string) string {

	return GetTime().Format(f)

}
