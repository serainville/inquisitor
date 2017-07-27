package cmd

import (
	"github.com/serainville/inquisitor/variables"
	"github.com/spf13/cobra"
)

// RootCmd is the CLIs root
var RootCmd = &cobra.Command{
	Use:   "inquisitor",
	Short: "System metrics collector",
	Long: `

  ▄████  ▒█████   ██▓     ██▓    ▓█████  ▄████▄  ▄▄▄█████▓ ▒█████   ██▀███  
 ██▒ ▀█▒▒██▒  ██▒▓██▒    ▓██▒    ▓█   ▀ ▒██▀ ▀█  ▓  ██▒ ▓▒▒██▒  ██▒▓██ ▒ ██▒
▒██░▄▄▄░▒██░  ██▒▒██░    ▒██░    ▒███   ▒▓█    ▄ ▒ ▓██░ ▒░▒██░  ██▒▓██ ░▄█ ▒
░▓█  ██▓▒██   ██░▒██░    ▒██░    ▒▓█  ▄ ▒▓▓▄ ▄██▒░ ▓██▓ ░ ▒██   ██░▒██▀▀█▄  
░▒▓███▀▒░ ████▓▒░░██████▒░██████▒░▒████▒▒ ▓███▀ ░  ▒██▒ ░ ░ ████▓▒░░██▓ ▒██▒
 ░▒   ▒ ░ ▒░▒░▒░ ░ ▒░▓  ░░ ▒░▓  ░░░ ▒░ ░░ ░▒ ▒  ░  ▒ ░░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░
  ░   ░   ░ ▒ ▒░ ░ ░ ▒  ░░ ░ ▒  ░ ░ ░  ░  ░  ▒       ░      ░ ▒ ▒░   ░▒ ░ ▒░
░ ░   ░ ░ ░ ░ ▒    ░ ░     ░ ░      ░   ░          ░      ░ ░ ░ ▒    ░░   ░ 
      ░     ░ ░      ░  ░    ░  ░   ░  ░░ ░  VER ` + variables.Version + `        ░ ░     ░     
                                        ░                                                                            
`,
}

func init() {

}
