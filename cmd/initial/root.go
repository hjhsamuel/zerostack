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
	initCmdFlags.StringVarP(&VarStringLogLevel, "level", "l", "info", "The default log level")
	initCmdFlags.IntVarP(&VarIntLogRolls, "rolls", "s", 3, "The max number of history log files")
	initCmdFlags.StringVarP(&VarStringHost, "host", "h", "0.0.0.0", "The host to listen on")
	initCmdFlags.IntVarP(&VarIntPort, "port", "p", 8080, "The port to listen on")
	initCmdFlags.StringVarP(&VarStringJwtSalt, "salt", "j", "zerostack", "The salt to generate jwt")
}
