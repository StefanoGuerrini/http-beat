package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stefanoguerrini/http-beat/core"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start enpoint monitoring.",
	Run: func(cmd *cobra.Command, args []string) {
		core.Monitor()
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
