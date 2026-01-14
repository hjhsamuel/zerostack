package file

import (
	"errors"
	"os"
	"path/filepath"
)

func CreateFileIfNotExist(path string) (*os.File, error) {
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

func MkdirIfNotExist(path string) error {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func WriteFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
