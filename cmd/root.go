package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "This is a test cli",
	Long: `This is a test cli long description
		to help explain what this command does and doesn't do'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a test")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
