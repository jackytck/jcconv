package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
	"github.com/jackytck/jcconv/web"
	"github.com/spf13/cobra"
)

var port = 8080

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// a. setup detector
		det, err := detector.NewDetector(0)
		if err != nil {
			log.Println(err)
			return
		}

		trans2hk, err := translator.New("s2hk")
		if err != nil {
			log.Println(err)
			return
		}
		trans2s, err := translator.New("hk2s")
		if err != nil {
			log.Println(err)
			return
		}

		// c. handlers
		log.Printf("Listening at http://127.0.0.1:%d...\n", port)
		http.HandleFunc("/", web.HelloServer)
		http.HandleFunc("/translate", web.Translate(det, trans2hk, trans2s))
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().IntVarP(&port, "port", "p", port, "Port of local server.")
}
