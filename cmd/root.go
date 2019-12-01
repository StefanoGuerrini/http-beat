package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http-beat",
	Short: "http-beat is a tool for check endpoints health.",
}

func Execute() {
	rootCmd.Execute()
}
