package cmdx

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func Check(name string) bool {
	_, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	return true
}

func Run(dir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
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
