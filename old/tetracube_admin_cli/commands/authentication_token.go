package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"tetracube.red/admin-cli/tetracube_admin_cli/services"
)

var authenticationTokenCmd = &cobra.Command{
	Use:     "auth-token",
	Aliases: []string{},
	Short:   "Authentication Token management",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		houseName := args[0]
		fmt.Printf("Creating authentication token for house %s on %s host\n", houseName, hostAddress)
		responseMap := services.CreateTokenForNewUser(hostAddress, houseName)
		if responseMap == nil {
			return
		}
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
