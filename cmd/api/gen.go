package api

import (
	"errors"

	"github.com/hjhsamuel/zerostack/cmd/pkg"
	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/gen/render"
	"github.com/hjhsamuel/zerostack/pkg/cmdx"
	"github.com/hjhsamuel/zerostack/pkg/file"
	"github.com/spf13/cobra"
)

var (
	// VarStringApiPath the api template path
	VarStringApiPath string
)

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
	if err = pkg.CheckProject(base, workspace); err != nil {
		return err
	}

	// generate api
	if err = render.CreateApiFile(base, VarStringApiPath); err != nil {
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
	if err = pkg.CheckProject(base, workspace); err != nil {
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
