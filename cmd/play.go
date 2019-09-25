package cmd

import (
	"fmt"
	"log"

	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
	"github.com/spf13/cobra"
)

var input string
var inPath string
var outPath string
var chain = "auto"
var thread = -1

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Translate a single line",
	Long:  "Enter original text in standard input and get back the result in standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		// a. detect chain
		if chain == "auto" {
			d, err := detector.NewDetector(0)
			if err != nil {
				log.Println(err)
				return
			}
			isTrad, err := d.IsTraditional(input)
			if err != nil {
				log.Println(err)
				return
			}
			if isTrad {
				chain = "hk2s"
			} else {
				chain = "s2hk"
			}
		}

		// b. translator
		trans, err := translator.New(chain)
		if err != nil {
			log.Println(err)
			return
		}

		// c. translate
		out, err := trans.Translate(input)
		must(err)

		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
	playCmd.Flags().StringVarP(&chain, "convert", "c", chain, "Conversion: one of 'auto' (default), 's2hk', 's2tw', 'hk2s', 'tw2s'.")
	must(playCmd.MarkFlagRequired("input"))
}
