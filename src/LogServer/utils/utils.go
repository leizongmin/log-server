package utils

import "os"

func Mkdirp(dir string) error {

	err := os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil

}
