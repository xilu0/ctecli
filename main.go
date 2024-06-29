package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xilu0/ctecli/cmd"
)

// init cobra
var rootCmd = &cobra.Command{
	Use:   "ctecli",
	Short: "coze cli",
	Long:  `coze cli`,
	Run: func(c *cobra.Command, args []string) {
		// Do Stuff Here
		var content string
		content, _ = c.Flags().GetString("content")
		if err := cmd.Call(content); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmd.ConfigCmd)
	rootCmd.Flags().StringP("content", "c", "", "content")

}
