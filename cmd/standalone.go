package cmd

import (
	"github.com/spf13/cobra"

	"github.com/serainville/inquisitor/models"
	"github.com/serainville/inquisitor/server"
)

var standaloneConfig models.ServerConfig

var standaloneCmd = &cobra.Command{
	Use:   "standalone",
	Short: "Run as standalone service",
	Run: func(cmd *cobra.Command, args []string) {
		//init := server.Init(config.IP, config.Port, TLSCertFile, TLSKeyFile, UseTLS)
		standaloneConfig.Standalone = true
		server.StartStandalone()

	},
}

func init() {

	RootCmd.AddCommand(standaloneCmd)
	standaloneCmd.Flags().BoolVarP(&standaloneConfig.Daemon, "daemon", "d", false, "Run server as daemon (service)")
}
