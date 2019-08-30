package cmd

import (
	"log"
	"time"

	"github.com/jackytck/go-chinese-converter/file"
	"github.com/jackytck/go-chinese-converter/translator"
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Translate a single file",
	Long:  "Translate a single file and output to the given path.",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		defer func() {
			if verbose {
				elapsed := time.Since(start)
				log.Println("Took", elapsed)
			}
		}()

		// a. check
		isText, err := file.IsPotentialTextFile(inPath)
		must(err)
		if !isText {
			log.Printf("%q is not a plain text file!\n", inPath)
			return
		}

		// b. translator
		trans, err := translator.New(chain)
		must(err)

		// c. translate
		must(trans.Translate(inPath, outPath))
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.Flags().StringVarP(&inPath, "file", "f", inPath, "Input file path.")
	fileCmd.Flags().StringVarP(&outPath, "out", "o", outPath, "Output file path.")
	fileCmd.Flags().IntVarP(&thread, "thread", "n", thread, "Number of threads to process, default is number of cores x 4")
	fileCmd.Flags().BoolVarP(&verbose, "verbose", "v", verbose, "Display more info")
	fileCmd.MarkFlagRequired("file")
}
