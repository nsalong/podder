package cmd

import (
	"github.com/spf13/cobra"
)

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Lists all pods according to the given --context=<NAME>",
	Long: `Lists all pods according to the given --context=<NAME>, showing detailed information about
		the pods within the given context`,
	Run: func(cmd *cobra.Command, args []string) {
		Handler.HandlePods(Context, OverridePath)
	},
}
