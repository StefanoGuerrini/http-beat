package cmd

import (
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start enpoint monitoring.",
	Run: func(cmd *cobra.Command, args []string) {
		Monitor()
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
