package api

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "api",
		Short: "Generate api code",
		Long:  "Generate api code",
		RunE:  GenerateApi,
	}
	swaggerCmd = &cobra.Command{
		Use:  "swagger",
		Long: "Generate swagger docs",
		RunE: GenerateSwagger,
	}
)

func init() {
	var (
		apiCmdFlags = Cmd.Flags()
	)

	apiCmdFlags.StringVarP(&VarStringApiPath, "path", "p", "api", "The api template path")

	Cmd.AddCommand(swaggerCmd)
}
