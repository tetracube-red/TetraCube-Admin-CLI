package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hostAddress string

var rootCmd = &cobra.Command{
	Use:   "manage",
	Short: "TetraCube Hub Manager CLI",
	Long:  "A Fast and Flexible CLI that helps user to initialize house just installed and create account tokens",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&hostAddress, "address", "", "Specify TetraCube's host address")
	_ = rootCmd.MarkPersistentFlagRequired("address")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
