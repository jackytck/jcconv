package web

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
)

// Index renders the index page with source represented as page.
func Index(page string, det *detector.Detector, tm map[string]*translator.Translator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// get input from url param if any
		var text, output, error string
		pageWithData := page
		var fromLocale, toLocale string
		switch r.Method {
		case http.MethodGet:
			texts, ok := r.URL.Query()["text"]
			if ok && len(texts[0]) > 0 {
				text = texts[0]

				// translate in ssr
				locale, err := det.DetectLang(text)
				if err != nil {
					error = err.Error()
					break
				}
				fromLocale = locale

				var trans *translator.Translator
				switch fromLocale {
				case translator.ZhHK:
					trans = tm["hk2s"]
					toLocale = translator.ZhCN
				case translator.ZhTW:
					trans = tm["tw2s"]
					toLocale = translator.ZhCN
				case translator.ZhCN:
					trans = tm["s2hk"]
					toLocale = translator.ZhHK
				}
				output, err = trans.Translate(text)
				if err != nil {
					error = err.Error()
					break
				}

				pageWithData = strings.Replace(pageWithData, " autofocus", "", 1)
			}
		}
		pageWithData = strings.Replace(pageWithData, "{INPUT}", text, 1)
		pageWithData = strings.Replace(pageWithData, "{OUTPUT}", output, 1)
		pageWithData = strings.Replace(pageWithData, "{ERROR}", error, 1)
		pageWithData = strings.Replace(pageWithData, "{FROM_LOCALE}", translator.LocaleToChineseText(fromLocale), 1)
		pageWithData = strings.Replace(pageWithData, "{TO_LOCALE}", translator.LocaleToChineseText(toLocale), 1)

		elapsed := time.Since(start).String()
		if text == "" {
			elapsed = ""
		}
		pageWithData = strings.Replace(pageWithData, "{ELAPSED}", elapsed, 1)

		fmt.Fprint(w, pageWithData)
	}
}
