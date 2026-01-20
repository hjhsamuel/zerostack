package render

import (
	_ "embed"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

//go:embed app.gotmpl
var appTpl string

const (
	appFilePath = "app/app.go"
)

func CreateAppFile(base *entities.BaseInfo, database string) error {
	absAppPath := filepath.Join(base.SrvHome, mainFilePath)
	if file.Exists(absAppPath) {
		return nil
	}
	if err := file.MkdirIfNotExist(filepath.Dir(absAppPath)); err != nil {
		return err
	}
	content, err := GetRenderedContentByParams("app", appTpl, map[string]any{
		"module": base.Module,
		"dao":    database,
	})
	if err != nil {
		return err
	}
	return file.WriteFile(absAppPath, content)
}
