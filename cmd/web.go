package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackytck/jcconv/box"
	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
	"github.com/jackytck/jcconv/web"
	"github.com/spf13/cobra"
)

var port = 8080
var domain = "http://127.0.0.1:8080"
var track = ""

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

		// b. setup translators in both directions (hk + tw)
		tm, err := translator.NewAll()
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
		domain = strings.Replace(domain, "8080", fmt.Sprintf("%d", port), 1)
		ps := strings.Replace(string(page), "{DOMAIN}", domain, 1)
		ps = strings.Replace(ps, "{ANALYTICS}", web.GoogleAnalytics(track), 1)

		// c. handlers
		log.Printf("Locally listening at http://127.0.0.1:%d\n", port)
		if !strings.Contains(domain, "127.0.0.1") {
			log.Printf("Externally at %s\n", domain)
		}
		http.HandleFunc("/", web.Index(ps, det, tm))
		http.HandleFunc("/translate", web.Translate(det, tm))
		must(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().IntVarP(&port, "port", "p", port, "Port of local server.")
	webCmd.Flags().StringVarP(&domain, "domain", "d", domain, "External protocol and domain name, e.g. https://jcconv.nat.com")
	webCmd.Flags().StringVarP(&track, "track", "t", track, "Google analytics tracking code, e.g. UA-XXXXXXXXX-X")
}
