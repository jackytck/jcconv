package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version info.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
