package pkg

import (
	"errors"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/gen/gomod"
)

func CheckProject(base *entities.BaseInfo, path string) error {
	// check go.mod
	isGoMod, modPath, err := gomod.IsGoMod(path)
	if err != nil {
		return err
	}
	if !isGoMod {
		return errors.New("missing go.mod, execute in project directory")
	}
	f, err := gomod.ParseModfile(modPath)
	if err != nil {
		return err
	}
	// check project
	service, err := gomod.GetServiceName(f.Module.Mod.Path)
	if err != nil {
		return err
	}
	if filepath.Base(path) != service {
		return errors.New("invalid project workspace")
	}

	base.Module = f.Module.Mod.Path
	base.Service = service
	return nil
}
