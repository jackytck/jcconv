package web

import "fmt"

// GoogleAnalytics returns the snippet of global site tag to be put into the
// first item of <head>.
func GoogleAnalytics(trackingCode string) string {
	if trackingCode == "" {
		return ""
	}
	var snippet = `
    <script async src="https://www.googletagmanager.com/gtag/js?id=%s"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());
      gtag('config', '%s');
    </script>`
	return fmt.Sprintf(snippet, trackingCode, trackingCode)
}
