package cmd

import (
	"fmt"
	"runtime"

	"github.com/hjhsamuel/zerostack/cmd/api"
	"github.com/hjhsamuel/zerostack/cmd/initial"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "zs",
		Long: "A CLI tool to generate golang microservice based on gin",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}
)

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", BuildVersion, runtime.GOOS, runtime.GOARCH)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.AddCommand(initial.Cmd, api.Cmd)
}

func Execute() error {
	return rootCmd.Execute()
}
