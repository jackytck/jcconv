package cmd

import (
	"fmt"
	"strings"

	"github.com/jackytck/go-chinese-converter/file"
	"github.com/jackytck/go-chinese-converter/lib"
	"github.com/spf13/cobra"
)

var input string
var inPath string
var outPath string

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long:  "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		main := []string{
			"STPhrases.txt",
			"STCharacters.txt",
		}
		variant := []string{
			"HKVariants.txt",
			"HKVariantsPhrases.txt",
			"HKVariantsRevPhrases.txt",
		}
		c, err := lib.NewChain(main, variant)
		if err != nil {
			panic(err)
		}
		if input != "" {
			out, err := c.Translate(input)
			if err != nil {
				panic(err)
			}
			fmt.Println(out)
		}

		if inPath != "" {
			s, err := file.ReadFile(inPath)
			if err != nil {
				panic(err)
			}

			var w []string
			for i, line := range s {
				t, err := c.Translate(line)
				if err != nil {
					panic(err)
				}
				w = append(w, t)
				if outPath == "" {
					fmt.Println(i+1, t)
				}
			}

			if outPath != "" {
				err := file.WriteFile(strings.Join(w, "\n"), outPath)
				if err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
	playCmd.Flags().StringVarP(&inPath, "file", "f", inPath, "Input file path.")
	playCmd.Flags().StringVarP(&outPath, "out", "o", outPath, "Output file path.")
}
