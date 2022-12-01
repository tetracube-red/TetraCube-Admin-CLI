package commands

import (
	"encoding/json"
	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
	"os"
	"tetracube.red/admin-cli/tetracube_admin_cli/services"
)

var authenticationTokenCmd = &cobra.Command{
	Use:     "auth-token",
	Aliases: []string{},
	Short:   "Authentication Token management",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		responseMap := services.CreateTokenForNewUser(hostAddress)
		responseMap["host"] = hostAddress
		jsonValue, _ := json.Marshal(responseMap)

		config := qrterminal.Config{
			Level:     qrterminal.H,
			Writer:    os.Stdout,
			BlackChar: qrterminal.WHITE,
			WhiteChar: qrterminal.BLACK,
			QuietZone: 1,
		}
		qrterminal.GenerateWithConfig(string(jsonValue), config)
	},
}

func init() {
	rootCmd.AddCommand(authenticationTokenCmd)
}
