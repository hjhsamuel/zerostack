package render

import (
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

const (
	swaggerDocsTpl = `//go:build swagger
// +build swagger

package docs

import "github.com/swaggo/swag"
`
	swaggerEmptyTpl = `//go:build !swagger

package docs
`
)

func CreateDocs(base *entities.BaseInfo) error {
	dir := filepath.Join(base.SrvHome, "docs")
	// create directory
	if err := file.MkdirIfNotExist(dir); err != nil {
		return err
	}
	// write file
	if err := file.WriteFile(filepath.Join(dir, "docs.go"), swaggerDocsTpl); err != nil {
		return err
	}
	if err := file.WriteFile(filepath.Join(dir, "empty.go"), swaggerEmptyTpl); err != nil {
		return err
	}

	return nil
}
