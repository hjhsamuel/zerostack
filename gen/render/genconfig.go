package render

import (
	_ "embed"
	"encoding/json"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

//go:embed config.gotmpl
var configTpl string

const (
	configFilePath = "config/config.go"
)

func CreateConfigFile(base *entities.BaseInfo, info *entities.ConfigInfo) error {
	absConfigPath := filepath.Join(base.SrvHome, configFilePath)
	// file exists, skip generation
	if file.Exists(absConfigPath) {
		return nil
	}

	if err := file.MkdirIfNotExist(filepath.Dir(absConfigPath)); err != nil {
		return err
	}

	content, err := json.Marshal(info)
	if err != nil {
		return err
	}
	var params map[string]any
	if err = json.Unmarshal(content, &params); err != nil {
		return err
	}
	params["module"] = base.Module
	params["service"] = base.Service

	code, err := GetRenderedContentByParams("config", configTpl, params)
	if err != nil {
		return err
	}
	return file.WriteFile(absConfigPath, code)
}
