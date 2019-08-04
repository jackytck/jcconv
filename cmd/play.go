package cmd

import (
	"github.com/jackytck/go-chinese-converter/dict"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long:  "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		chains := []string{
			"HKVariants.txt",
			"TWVariants.txt",
		}
		_, err := dict.NewDict(chains)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
