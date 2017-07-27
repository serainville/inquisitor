package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	Host string
	Port int
	Cert string
	Tls  bool
)

func init() {
	RootCmd.AddCommand(agentCmd)
	agentCmd.Flags().StringVarP(&Host, "host", "s", "localhost", "Gollector server name or IP")
	agentCmd.Flags().IntVarP(&Port, "port", "p", 21200, "Gollector server port number")
	agentCmd.Flags().StringVarP(&Cert, "cert", "", "", "Agent certificate file")
	agentCmd.Flags().BoolVarP(&Tls, "tls", "", false, "Use TLS")
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Agent management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connecting to server....")
		fmt.Println(" Server: " + Host)
		fmt.Println(" Port: " + strconv.Itoa(Port))
	},
}
