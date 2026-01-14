package file

import "os"

func GetWorkspace() (string, error) {
	return os.Getwd()
}
