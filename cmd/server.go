package cmd

import (
	"github.com/spf13/cobra"

	"github.com/serainville/inquisitor/models"
	"github.com/serainville/inquisitor/server"
)

var (
	config models.ServerConfig
)

func init() {

	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&config.IP, "ip", "b", "127.0.0.1", "Bind server to IP address")
	serverCmd.Flags().StringVarP(&config.Port, "port", "p", "27001", "Bind server to port")
	serverCmd.Flags().BoolVarP(&config.Daemon, "daemon", "d", false, "Run server as daemon (service)")
	serverCmd.Flags().BoolVarP(&config.UseTLS, "usetls", "", false, "Enable TLS")
	serverCmd.Flags().StringVarP(&config.TLSCertFile, "cert", "", "", "TLS certificate file")
	serverCmd.Flags().StringVarP(&config.TLSKeyFile, "key", "", "", "TLS key file")

}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Daemon management",
	Run: func(cmd *cobra.Command, args []string) {
		//init := server.Init(config.IP, config.Port, TLSCertFile, TLSKeyFile, UseTLS)
		server.StartServer(&config)
	},
}
