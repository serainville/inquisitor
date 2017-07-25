package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"inquisitor/constants"
)


func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.AppName + " " + constants.Version)
	},
}


