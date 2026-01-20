package render

import (
	_ "embed"
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

	if err := file.MkdirIfNotExist(filepath.Dir(absConfigPath)); err != nil {
		return err
	}
	return file.WriteFile(absConfigPath, configTpl)
}
