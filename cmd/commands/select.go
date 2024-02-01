// select.go
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a subscription",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		selectSubscription(args[0])
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}

func selectSubscription(subscriptionID string) {
	fmt.Printf("Selected subscription with ID: %s\n", subscriptionID)
}
