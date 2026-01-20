package render

import (
	_ "embed"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

//go:embed server.gotmpl
var serverTpl string

const (
	serverFilePath = "internal/server/server.go"
)

func CreateServerFile(base *entities.BaseInfo) error {
	path := filepath.Join(base.SrvHome, serverFilePath)
	if file.Exists(path) {
		return nil
	}

	return CreateGoTemplate(serverTpl, path, &entities.GenInfo{
		Module:  base.Module,
		Service: base.Service,
	})
}
