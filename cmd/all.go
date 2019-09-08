package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackytck/jcconv/file"
	"github.com/jackytck/jcconv/translator"
	"github.com/spf13/cobra"
)

var inDir, outDir string
var verbose bool

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Translate all text file(s)",
	Long:  "Translate all of the text file(s) under the input directory.",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		defer func() {
			if verbose {
				elapsed := time.Since(start)
				log.Println("Took", elapsed)
			}
		}()

		// check
		c, err := file.IsDir(inDir)
		must(err)
		if !c {
			fmt.Printf("%q is not a directory!", inDir)
		}
		must(file.EnsureDir(outDir, 0775))

		root, err := filepath.Abs(inDir)
		must(err)

		trans, err := translator.New(chain)
		must(err)

		onWalk := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			dst := strings.Replace(path, root, outDir, 1)

			if info.Mode().IsRegular() {
				text, err := file.IsPotentialTextFile(path)
				must(err)
				if text {
					// translate text
					must(trans.Translate(path, dst))
				} else {
					// just copy any file
					_, err := file.Copy(path, dst)
					must(err)
				}
			} else {
				must(file.EnsureDir(dst, 0775))
			}

			return nil
		}

		// walk
		err = filepath.Walk(inDir, onWalk)
		must(err)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().StringVarP(&inDir, "in", "d", inDir, "Input directory path.")
	allCmd.Flags().StringVarP(&outDir, "out", "o", outDir, "Output directory path.")
	allCmd.Flags().StringVarP(&chain, "convert", "c", chain, "Conversion: one of 's2hk' (default), 's2tw', 'hk2s', 'tw2s'.")
	allCmd.Flags().BoolVarP(&verbose, "verbose", "v", verbose, "Display more info")
	allCmd.MarkFlagRequired("in")
	allCmd.MarkFlagRequired("out")
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
