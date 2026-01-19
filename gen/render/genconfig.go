package render

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

//go:embed config.gotmpl
var configTpl string

const (
	configFilePath = "config/config.go"
)

func CreateConfigFile(base *entities.BaseInfo) error {
	absConfigPath := filepath.Join(base.SrvHome, configFilePath)
	// file exists, skip generation
	if file.Exists(absConfigPath) {
		return nil
	}

	f, err := os.Create(absConfigPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(configTpl)
	return err
}
