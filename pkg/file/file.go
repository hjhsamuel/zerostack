package file

import (
	"errors"
	"os"
	"path/filepath"
)

var Workspace string

func init() {
	dir, err := os.Getwd()
	if err == nil {
		Workspace = dir
	}
}

func CreateIfNotExist(path string) (*os.File, error) {
	if Exists(path) {
		return nil, errors.New("file already exists")
	}
	parent := filepath.Dir(path)
	if !Exists(parent) {
		if err := os.MkdirAll(parent, os.ModePerm); err != nil {
			return nil, err
		}
	}
	return os.Create(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
