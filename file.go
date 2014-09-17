package utils

import "os"

func FileExists(name string) (bool, error) {
	finfo, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	if finfo.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}
