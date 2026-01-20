package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/gen/gomod"
	"github.com/hjhsamuel/zerostack/gen/render"
	"github.com/hjhsamuel/zerostack/pkg/cmdx"
	"github.com/hjhsamuel/zerostack/pkg/file"
	"github.com/spf13/cobra"
)

func InitService(_ *cobra.Command, _ []string) error {
	// check module
	if VarStringModule == "" {
		return errors.New("missing -module")
	}
	module, err := gomod.ParseModuleName(VarStringModule)
	if err != nil {
		return fmt.Errorf("invalid module [%s]: %s", VarStringModule, err.Error())
	}
	service, err := gomod.GetServiceName(module)
	if err != nil {
		return err
	}
	// get workspace
	workspace, err := file.GetWorkspace()
	if err != nil {
		return err
	}

	base := &entities.BaseInfo{
		Module:    module,
		Service:   service,
		Workspace: workspace,
		SrvHome:   filepath.Join(workspace, service),
	}
	if file.Exists(base.SrvHome) {
		return errors.New("service already exists")
	}

	// generate main
	if err = render.CreateMainFile(base); err != nil {
		goto Err
	}
	// generate config
	if err = render.CreateConfigFile(base); err != nil {
		goto Err
	}
	// generate database
	if err = render.CreateDaoFile(base, VarStringDB); err != nil {
		goto Err
	}
	// generate app
	if err = render.CreateAppFile(base, VarStringDB); err != nil {
		goto Err
	}
	// generate api
	if err = render.CreateApiFile(base, VarStringApiFile); err != nil {
		goto Err
	}
	// generate docs
	if err = render.CreateDocs(base); err != nil {
		goto Err
	}
	// generate pkg
	if err = render.CreatePkg(base); err != nil {
		goto Err
	}
	// prepare go.mod
	if err = render.PrepareGoMod(base); err != nil {
		goto Err
	}

	return nil

Err:
	_ = file.RemoveAll(base.SrvHome)
	return err
}

func checkProject(base *entities.BaseInfo, path string) error {
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

func GenerateApi(_ *cobra.Command, _ []string) error {
	// get workspace
	workspace, err := file.GetWorkspace()
	if err != nil {
		return err
	}

	base := &entities.BaseInfo{
		Workspace: workspace,
		SrvHome:   workspace,
	}
	if err = checkProject(base, workspace); err != nil {
		return err
	}

	// generate api
	if err = render.CreateApiFile(base, VarStringApiFile); err != nil {
		return err
	}

	return nil
}

func GenerateSwagger(_ *cobra.Command, _ []string) error {
	// get workspace
	workspace, err := file.GetWorkspace()
	if err != nil {
		return err
	}

	base := &entities.BaseInfo{
		Workspace: workspace,
		SrvHome:   workspace,
	}
	if err = checkProject(base, workspace); err != nil {
		return err
	}

	// generate swagger
	if !cmdx.Check("swag") {
		return errors.New("invalid command: swag")
	}
	_, err = cmdx.Run(base.SrvHome, "swag", "init", "--pd", "-g", "./cmd/main.go")
	if err != nil {
		return err
	}

	return nil
}
