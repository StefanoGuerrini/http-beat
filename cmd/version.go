package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show http-beat version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http-beat v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
