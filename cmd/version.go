package cmd

import (
	"fmt"

	"github.com/serainville/inquisitor/variables"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(variables.AppName + " " + variables.Version)
	},
}
