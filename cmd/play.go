package cmd

import (
	"fmt"

	"github.com/jackytck/go-chinese-converter/lib"
	"github.com/spf13/cobra"
)

var input string

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long:  "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		chains := []string{
			"STPhrases.txt",
			"STCharacters.txt",
			"HKVariants.txt",
			"HKVariantsPhrases.txt",
			"HKVariantsRevPhrases.txt",
		}
		c, err := lib.NewChain(chains)
		if err != nil {
			panic(err)
		}
		out, err := c.Translate(input)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
}
