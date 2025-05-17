package cmd

import (
	"crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash [text]",
	Short: "Hash a string using SHA256",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		hash := sha256.Sum256([]byte(input))
		fmt.Printf("SHA256(%s): %x\n", input, hash)
	},
}
