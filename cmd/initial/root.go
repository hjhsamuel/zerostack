package initial

import "github.com/spf13/cobra"

var (
	Cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long:  "Initialize a new project",
		RunE:  InitProject,
	}
)

func init() {
	var (
		initCmdFlags = Cmd.Flags()
	)

	initCmdFlags.StringVarP(&VarStringModule, "module", "m", "", "The module name used in go.mod")
	initCmdFlags.StringVarP(&VarStringDatabase, "database", "t", "mysql", "The database type")
	initCmdFlags.StringVarP(&VarStringApiPath, "path", "p", "api", "The api template path")
}
