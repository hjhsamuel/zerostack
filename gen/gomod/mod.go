package gomod

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

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

func GetServiceName(name string) (string, error) {
	baseName := filepath.Base(name)
	re := regexp.MustCompile(`^(?P<name>[a-z]+[\w\-]*).*$`)
	rsp := re.FindStringSubmatch(baseName)
	for index, item := range re.SubexpNames() {
		if index == 0 || item == "" {
			continue
		}
		return rsp[index], nil
	}
	return "", errors.New("service name not found")
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
