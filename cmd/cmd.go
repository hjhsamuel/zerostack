package cmd

import (
	"errors"

	"github.com/hjhsamuel/zerostack/pkg/gomod"
	"github.com/spf13/cobra"
)

func InitService(_ *cobra.Command, _ []string) error {
	moduleName := VarStringModule
	if moduleName == "" {
		return errors.New("missing command: -module")
	}

	var err error
	moduleName, err = gomod.ParseModuleName(moduleName)
	if err != nil {
		return err
	}
}
