package render

import (
	_ "embed"
	"os"
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
	content, err := GetRenderedContentByParams("app", appTpl, map[string]any{
		"module": base.Module,
		"dao":    database,
	})
	if err != nil {
		return err
	}
	f, err := os.Create(absAppPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}
