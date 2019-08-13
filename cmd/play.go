package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jackytck/go-chinese-converter/file"
	"github.com/jackytck/go-chinese-converter/lib"
	"github.com/spf13/cobra"
)

var input string
var inPath string
var outPath string
var thread = -1

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
			isText, err := file.IsTextFile(inPath)
			if err != nil {
				panic(err)
			}
			if !isText {
				log.Printf("%q is not a plain text file!\n", inPath)
				return
			}

			lines, size, errc := file.ScanFile(inPath)
			done := make(chan struct{})
			// defer close(done)
			result := make(chan file.Line)

			digester := file.Digester{
				Chain:  c,
				Done:   done,
				Lines:  lines,
				Result: result,
			}
			digester.Run(-1)

			w := make([]string, size)
			for l := range result {
				w[l.LineNum] = l.Text
			}
			if err := <-errc; err != nil {
				panic(err)
			}

			if outPath != "" {
				err := file.WriteFile(strings.Join(w, "\n"), outPath)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println(strings.Join(w, "\n"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
	playCmd.Flags().StringVarP(&inPath, "file", "f", inPath, "Input file path.")
	playCmd.Flags().StringVarP(&outPath, "out", "o", outPath, "Output file path.")
	playCmd.Flags().IntVarP(&thread, "thread", "n", thread, "Number of threads to process, default is number of cores x 4")
}
