// root.go
package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azurectl",
	Short: "Azure subscription manager",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
