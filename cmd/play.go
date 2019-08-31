package cmd

import (
	"fmt"
	"log"

	"github.com/jackytck/go-chinese-converter/translator"
	"github.com/spf13/cobra"
)

var input string
var inPath string
var outPath string
var chain = "s2hk"
var thread = -1

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Translate a single line",
	Long:  "Enter original text in standard input and get back the result in standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		// a. translator
		trans, err := translator.New(chain)
		if err != nil {
			log.Println(err)
			return
		}

		// b. translate
		out, err := trans.TranslateOne(input)
		must(err)

		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
	playCmd.Flags().StringVarP(&chain, "convert", "c", chain, "Conversion: one of 's2hk' (default), 's2tw', 'hk2s', 'tw2s'.")
	playCmd.MarkFlagRequired("input")
}
