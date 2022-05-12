package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"podder/handler"
)

var Handler *handler.CmdHandler

var Context string
var OverridePath string

var podderCmd = &cobra.Command{
	Use:   "podder",
	Short: "Podder is used to get detailed pod info",
	Long:  `Podder is used to get detailed pod information from kubernetes`,
	Run: func(cmd *cobra.Command, args []string) {
		Handler.HandlePodder()
	},
}

func Execute() {
	if err := podderCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	Handler = &handler.CmdHandler{Finished: false}

	podderCmd.AddCommand(verifyCmd)

	podsCmd.Flags().StringVarP(&Context, "context", "c", "", "Context for k8s")
	podsCmd.Flags().StringVarP(&OverridePath, "overridePath", "p", "", "Manual k8s config override")
	podderCmd.AddCommand(podsCmd)

}
