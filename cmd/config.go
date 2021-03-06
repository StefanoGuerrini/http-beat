package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"http-beat/core"
	"os"
)

var (
	urls    []string
	seconds int

	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure http-beat.",
		Run: func(cmd *cobra.Command, args []string) {

			if cmd.Flags().NFlag() == 0 {
				cmd.Help()
				os.Exit(0)
			}

			seconds, _ := cmd.Flags().GetInt("seconds")
			if seconds > 0 {
				core.SetBeatSeconds(seconds)
			}

			urls, _ := cmd.Flags().GetStringArray("urls")
			if len(urls) > 0 {
				core.AddUrls(urls)
			}

			get, _ := cmd.Flags().GetBool("get")
			if get {
				fmt.Println(core.GetBeatSeconds())
				fmt.Println(core.GetUrls())
			}

		},
	}
)

func init() {
	configCmd.Flags().StringArrayVarP(&urls, "urls", "u", make([]string, 0), "add urls to check")
	configCmd.Flags().IntVarP(&seconds, "seconds", "s", 0, "set seconds")
	configCmd.Flags().BoolP("get", "g", false, "get configuration")

	rootCmd.AddCommand(configCmd)
}
