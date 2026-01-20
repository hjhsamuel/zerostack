package initial

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/gen/gomod"
	"github.com/hjhsamuel/zerostack/gen/render"
	"github.com/hjhsamuel/zerostack/pkg/file"
	"github.com/spf13/cobra"
)

var (
	// VarStringModule the module name used in go.mod
	VarStringModule string
	// VarStringDatabase the database type
	VarStringDatabase string
	// VarStringApiPath the api template path
	VarStringApiPath string
)

func InitProject(_ *cobra.Command, _ []string) error {
	// check module
	if VarStringModule == "" {
		return errors.New("missing -m/--module")
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
	if err = render.CreateDaoFile(base, VarStringDatabase); err != nil {
		goto Err
	}
	// generate app
	if err = render.CreateAppFile(base, VarStringDatabase); err != nil {
		goto Err
	}
	// generate api
	if err = render.CreateApiFile(base, VarStringApiPath); err != nil {
		goto Err
	}
	// generate server
	if err = render.CreateServerFile(base); err != nil {
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
