package cmd

import (
	"log"

	"github.com/jackytck/go-chinese-converter/box"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long:  "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		hkv, ok := box.Get("/HKVariants.txt")
		if !ok {
			log.Println("Not found!")
		}
		log.Println(string(hkv))
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
