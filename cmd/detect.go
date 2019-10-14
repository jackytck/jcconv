package cmd

import (
	"fmt"

	"github.com/jackytck/jcconv/detector"
	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detect if given string is traditional or simplified",
	Long:  "Detect if a given string is traditional or simplified. It is classified as traditional only if it has 97% confidence.",
	Run: func(cmd *cobra.Command, args []string) {
		d, err := detector.NewDetector(0)
		must(err)
		lang, err := d.DetectLang(input)
		must(err)
		fmt.Println(lang)
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
	detectCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
	must(detectCmd.MarkFlagRequired("input"))
}
