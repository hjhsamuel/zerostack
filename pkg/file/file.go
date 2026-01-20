package file

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

func IsSame(src, dst string) bool {
	cSrc := filepath.Clean(src)
	cDst := filepath.Clean(dst)
	if runtime.GOOS == "windows" {
		return strings.ToLower(cSrc) == strings.ToLower(cDst)
	}
	return cSrc == cDst
}

func IsSubPath(path, base string) bool {
	cPath := strings.TrimSuffix(filepath.Clean(path), string(filepath.Separator))
	cBase := strings.TrimSuffix(filepath.Clean(base), string(filepath.Separator))
	if runtime.GOOS == "windows" {
		cPath = strings.ToLower(cPath)
		cBase = strings.ToLower(cBase)
	}
	return strings.HasPrefix(cPath, cBase)
}

func Copy(src, dir string) (string, error) {
	info, err := os.Stat(src)
	if err != nil {
		return "", err
	}
	if err = MkdirIfNotExist(dir); err != nil {
		return "", err
	}
	if info.IsDir() {
		entries, err := os.ReadDir(src)
		if err != nil {
			return "", err
		}
		for _, entry := range entries {
			_, err = Copy(filepath.Join(src, entry.Name()), filepath.Join(dir, entry.Name()))
			if err != nil {
				return "", err
			}
		}
		return filepath.Join(dir, info.Name()), nil
	} else {
		path := filepath.Join(dir, info.Name())
		err = copyTo(src, path)
		if err != nil {
			return "", err
		}
		return path, nil
	}
}

func copyTo(src, dst string) error {
	if Exists(dst) {
		return nil
	}

	rf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer rf.Close()

	wf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer wf.Close()

	_, err = io.Copy(wf, rf)
	if err != nil {
		return err
	}
	return nil
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}
