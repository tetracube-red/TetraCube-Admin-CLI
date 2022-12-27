package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"tetracube.red/admin-cli/tetracube_admin_cli/services"
)

var houseCmd = &cobra.Command{
	Use:     "house",
	Aliases: []string{},
	Short:   "House management",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		houseName := args[0]
		fmt.Printf("Creating house %s on %s host\n", houseName, hostAddress)
		services.CreateHouse(hostAddress, houseName)
	},
}

func init() {
	rootCmd.AddCommand(houseCmd)
}
