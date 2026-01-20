package render

import (
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

func CreatePkg(base *entities.BaseInfo) error {
	dir := filepath.Join(base.SrvHome, "pkg")
	return file.MkdirIfNotExist(dir)
}
