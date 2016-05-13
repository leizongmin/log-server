// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package server

import "LogServer/utils"

func setDir(dir string) error {

	if err := utils.Mkdirp(dir); err != nil {
		return err
	}

	return nil

}
