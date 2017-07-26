package cmd

import (
	"github.com/spf13/cobra"

	"github.com/serainville/inquisitor/server"
)

var (
	BindIP string
	BindPort string
	Daemon bool
	UseTLS bool
	TLSCertFile string
	TLSKeyFile string

	DBHost string
	DBUser string
	DBUserHasPassword bool
)

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&BindIP, "ip","b","127.0.0.1", "Bind server to IP address")
	serverCmd.Flags().StringVarP(&BindPort, "port","p", "27001", "Bind server to port")
	serverCmd.Flags().BoolVarP(&Daemon, "daemon","d", false, "Run server as daemon (service)")
	serverCmd.Flags().BoolVarP(&UseTLS, "usetls","", false, "Enable TLS")
	serverCmd.Flags().StringVarP(&TLSCertFile, "cert","", "", "TLS certificate file")
	serverCmd.Flags().StringVarP(&TLSKeyFile, "key","", "", "TLS key file")

}

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Daemon management",
	Run: func(cmd *cobra.Command, args []string) {
		init := server.Init(BindIP, BindPort, TLSCertFile, TLSKeyFile, UseTLS)
		server.StartServer(init)
	},
}


