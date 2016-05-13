// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package utils

import "os"

// Mkdirp Like `os.MkdirAll()`, but if directory already exists don't returns error
func Mkdirp(dir string) error {

	err := os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil

}
