package cmd

import (
	"fmt"
	"os"

	ollama "github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

// summarizeCmd represents the summarize command
var summarizeCmd = &cobra.Command{
	Use:   "summarize <filename>",
	Short: "Summarize the code in the given filename",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := ollama.ClientFromEnvironment()
		if err != nil {
			return err
		}
		contents, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		req := &ollama.ChatRequest{
			Model: "granite-code",
			Messages: []ollama.Message{
				{
					Role:    "user",
					Content: "explain in a single sentence the purpose of the following python code: \n\n`" + string(contents),
				},
			},
		}
		err = cl.Chat(cmd.Context(), req, func(resp ollama.ChatResponse) error {
			if msg := resp.Message; msg.Content != "" {
				fmt.Fprint(cmd.OutOrStdout(), msg.Content)
			}
			return nil
		})
		return err
	},
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
}
