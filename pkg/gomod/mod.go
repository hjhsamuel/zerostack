package gomod

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"

	"strings"
)

func ParseModuleName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("module name is empty")
	}
	if err := module.CheckPath(name); err != nil {
		return "", err
	}
	return name, nil
}

func ParseModfile(path string) (*modfile.File, error) {
	if name := filepath.Base(path); name != "go.mod" {
		return nil, errors.New("not a go.mod file")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return modfile.Parse("go.mod", content, nil)
}
