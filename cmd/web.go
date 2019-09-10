package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackytck/jcconv/box"
	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
	"github.com/jackytck/jcconv/web"
	"github.com/spf13/cobra"
)

var port = 8080

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start translation server",
	Long:  "Start a translation api server, plus a web interface for interacting with it.",
	Run: func(cmd *cobra.Command, args []string) {

		// a. setup detector
		det, err := detector.NewDetector(0)
		if err != nil {
			log.Println(err)
			return
		}

		// b. setup translators in both directions
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

		// c. setup index page
		page, ok := box.Get("/index.html")
		if !ok {
			log.Println("index.html not found!")
			return
		}

		// c. handlers
		log.Printf("Listening at http://127.0.0.1:%d\n", port)
		http.HandleFunc("/", web.Index(string(page)))
		http.HandleFunc("/translate", web.Translate(det, trans2hk, trans2s))
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().IntVarP(&port, "port", "p", port, "Port of local server.")
}
