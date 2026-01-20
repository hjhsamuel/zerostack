package render

import (
	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/gen/gomod"
)

func PrepareGoMod(base *entities.BaseInfo) error {
	// go mod init
	_, err := gomod.PrepareGoModule(base.SrvHome, base.Module)
	if err != nil {
		return err
	}
	// go mod tidy
	// TODO
	return nil
}
