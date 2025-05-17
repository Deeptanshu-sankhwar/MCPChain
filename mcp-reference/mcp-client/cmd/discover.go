package cmd

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover MCP servers registered on-chain",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("mcpchaind", "query", "toolregistry", "all-servers", "--output", "json").Output()
		if err != nil {
			fmt.Println("Error querying chain:", err)
			return
		}

		var prettyJSON map[string]interface{}
		if err := json.Unmarshal(out, &prettyJSON); err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		prettyOut, _ := json.MarshalIndent(prettyJSON, "", "  ")
		fmt.Println(string(prettyOut))
	},
}
