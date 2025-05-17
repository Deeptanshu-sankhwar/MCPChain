package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	endpoint string
	payload  string
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invoke /echo or /time endpoint of an MCP server",
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

		fmt.Println("Response:", string(body))
		fmt.Println("Request SHA256:", fmt.Sprintf("%x", reqHash))
		fmt.Println("Response SHA256:", fmt.Sprintf("%x", respHash))

		fmt.Println("\n📦 JSON:")
		result := map[string]interface{}{
			"request":       payload,
			"response":      string(body),
			"request_hash":  fmt.Sprintf("%x", reqHash),
			"response_hash": fmt.Sprintf("%x", respHash),
			"timestamp":     time.Now().Unix(),
		}
		pretty, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(pretty))
	},
}

func init() {
	invokeCmd.Flags().StringVar(&endpoint, "endpoint", "", "MCP server endpoint URL")
	invokeCmd.Flags().StringVar(&payload, "payload", "", "Payload for POST request")
	invokeCmd.MarkFlagRequired("endpoint")
}
