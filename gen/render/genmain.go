package render

import (
	_ "embed"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

//go:embed main.gotmpl
var mainTpl string

const (
	mainFilePath = "cmd/main.go"
)

func CreateMainFile(base *entities.BaseInfo) error {
	absMainPath := filepath.Join(base.Home, base.Service, mainFilePath)
	if file.Exists(absMainPath) {
		return nil
	}
	return CreateGoTemplate(mainTpl, absMainPath, map[string]any{
		"module":  base.Module,
		"service": base.Service,
	})
}
