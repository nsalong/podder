package cmd

import (
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Attempts to find your kubernetes config",
	Long:  `Attempts to find your kubernetes config from a default location`,
	Run: func(cmd *cobra.Command, args []string) {
		Handler.HandleVerify()
	},
}
