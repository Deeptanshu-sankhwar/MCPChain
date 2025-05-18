package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var (
	endpoint   string
	payload    string
	agentID    string
	toolID     uint64
	toolName   string
	fromKey    string
	skipAttest bool
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invoke /echo or /time endpoint of an MCP server and log attestation",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{Timeout: 5 * time.Second}

		var req *http.Request
		var err error

		if payload != "" {
			req, err = http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(payload)))
			req.Header.Set("Content-Type", "application/json")
		} else {
			req, err = http.NewRequest("GET", endpoint, nil)
		}

		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error calling endpoint:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		reqHash := sha256.Sum256([]byte(payload))
		respHash := sha256.Sum256(body)
		timestamp := time.Now().Unix()

		fmt.Println("Response:", string(body))
		fmt.Println("Request SHA256:", fmt.Sprintf("%x", reqHash))
		fmt.Println("Response SHA256:", fmt.Sprintf("%x", respHash))

		fmt.Println("\nJSON:")
		result := map[string]interface{}{
			"request":       payload,
			"response":      string(body),
			"request_hash":  fmt.Sprintf("%x", reqHash),
			"response_hash": fmt.Sprintf("%x", respHash),
			"timestamp":     timestamp,
		}
		pretty, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(pretty))

		// 🧾 Submit attestation tx
		if !skipAttest {
			fmt.Println("\nSubmitting on-chain attestation...")

			cmd := exec.Command("mcpchaind", "tx", "attestation", "log-attestation",
				agentID,
				fmt.Sprint(toolID),
				toolName,
				fmt.Sprintf("%x", reqHash),
				fmt.Sprintf("%x", respHash),
				fmt.Sprint(timestamp),
				"--from", fromKey,
				"--chain-id", "mcpchain",
				"--keyring-backend", "test",
				"--gas", "auto",
				"--fees", "1000umcp",
				"-y",
			)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Failed to submit attestation:", err)
			}
		}
	},
}

func init() {
	invokeCmd.Flags().StringVar(&endpoint, "endpoint", "", "MCP server endpoint URL")
	invokeCmd.Flags().StringVar(&payload, "payload", "", "Payload for POST request")
	invokeCmd.Flags().StringVar(&agentID, "agent-id", "", "Agent ID for attestation")
	invokeCmd.Flags().Uint64Var(&toolID, "tool-id", 0, "Tool Server ID from MCPChain")
	invokeCmd.Flags().StringVar(&toolName, "tool-name", "", "Tool name (used in attestation)")
	invokeCmd.Flags().StringVar(&fromKey, "from", "", "CLI key name to sign transaction")
	invokeCmd.Flags().BoolVar(&skipAttest, "no-attest", false, "Skip attestation on-chain")

	invokeCmd.MarkFlagRequired("endpoint")
	invokeCmd.MarkFlagRequired("agent-id")
	invokeCmd.MarkFlagRequired("tool-id")
	invokeCmd.MarkFlagRequired("tool-name")
	invokeCmd.MarkFlagRequired("from")
}
