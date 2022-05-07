package path

import (
	"errors"
	"os"
	"path/filepath"
)

func MakeParentDir(path string) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
}

func PathIsExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func GetPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}
