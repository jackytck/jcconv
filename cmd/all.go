package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackytck/go-chinese-converter/file"
	"github.com/spf13/cobra"
)

var inDir, outDir string

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Translate all text file(s)",
	Long:  "Translate all of the text file(s) under the input directory.",
	Run: func(cmd *cobra.Command, args []string) {
		// check
		c, err := file.IsDir(inDir)
		must(err)
		if !c {
			fmt.Printf("%q is not a directory!", inDir)
		}
		must(file.EnsureDir(outDir, 0775))

		root, err := filepath.Abs(inDir)
		must(err)

		onWalk := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Mode().IsRegular() {
				//@TODO translate text file
				//@TODO copy non-text file
			} else {
				d := strings.Replace(path, root, outDir, 1)
				must(file.EnsureDir(d, 0775))
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
	allCmd.MarkFlagRequired("in")
	allCmd.MarkFlagRequired("out")
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
