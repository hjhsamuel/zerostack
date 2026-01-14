package gomod

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func InitService(dir, module string) (string, error) {
	if _, err := exec.LookPath("go"); err != nil {
		return "", err
	}
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = dir

	var (
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
	)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() != 0 {
			msg := strings.TrimSuffix(stderr.String(), "\n")
			return "", errors.New(msg)
		}
		return "", err
	}
	msg := strings.TrimSuffix(stdout.String(), "\n")
	return msg, nil
}
