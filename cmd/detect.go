package cmd

import (
	"fmt"

	"github.com/jackytck/go-chinese-converter/detector"
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
		t, err := d.IsTraditional(input)
		must(err)
		if t {
			fmt.Println("traditional")
		} else {
			fmt.Println("simplified")
		}
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
	detectCmd.Flags().StringVarP(&input, "input", "i", input, "Input string.")
	detectCmd.MarkFlagRequired("input")
}
