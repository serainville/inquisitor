package cmd

import (
	"github.com/spf13/cobra"

	"github.com/serainville/inquisitor/client"
)

var (
	// Host is the hostname or IP address of server
	Host string
	// Port number of server
	Port int
	// Cert is the certificate file name
	Cert string
	// TLS sets whether TLS support is Enabled or disable
	TLS bool
)

func init() {
	RootCmd.AddCommand(agentCmd)
	agentCmd.Flags().StringVarP(&Host, "host", "s", "localhost", "Gollector server name or IP")
	agentCmd.Flags().IntVarP(&Port, "port", "p", 21200, "Gollector server port number")
	agentCmd.Flags().StringVarP(&Cert, "cert", "", "", "Agent certificate file")
	agentCmd.Flags().BoolVarP(&TLS, "tls", "", false, "Use TLS")
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Agent management",
	Run: func(cmd *cobra.Command, args []string) {
		client.Start()
	},
}
