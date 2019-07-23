package main

import (
	"os"

	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.project_name}}/internal/version"

	"github.com/spf13/cobra"
)

// versionCmd returns a CLI command that when run prints
// the application build version, commit and time
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the build version",
		Run: func(*cobra.Command, []string) {
			version.Write(os.Stdout)
		},
	}
}
