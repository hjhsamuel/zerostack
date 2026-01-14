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
	absConfigPath := filepath.Join(base.Home, base.Service, configFilePath)
	// file exists, skip generation
	if file.Exists(absConfigPath) {
		return nil
	}

	return CreateGoTemplate(configTpl, absConfigPath, map[string]any{})
}
