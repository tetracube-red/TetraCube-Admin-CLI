package house_admin_cli

import (
	"github.com/spf13/cobra"
	// "tetracube.red/admin-cli/tetracube_admin_cli/admin_cli_root"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//res := admin_cli_root.Reverse(args[0])
		//fmt.Println(res)
	},
}

func init() {
	// rootCmd.AddCommand(reverseCmd)
}
