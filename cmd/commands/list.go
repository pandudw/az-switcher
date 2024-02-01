// list.go
package commands

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display subscription list",
	Run: func(cmd *cobra.Command, args []string) {
		listSubscriptions()
	},
}

func listSubscriptions() {
	cmd := exec.Command("az", "account", "list", "--output", "yaml")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error executing Azure CLI command: %v\n", err)
		return
	}

	fmt.Printf("\n%s\n", output)
}
