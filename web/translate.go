package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
)

// Translate auto translates the input text.
func Translate(det *detector.Detector, tm map[string]*translator.Translator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		var res TransRes
		defer func() {
			res.Elapsed = time.Since(start).String()
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				panic(err)
			}
		}()

		var text string
		var chain string
		switch r.Method {
		case http.MethodGet:
			q := r.URL.Query()
			texts, ok := q["text"]
			if !ok || len(texts[0]) < 1 {
				res.Error = fmt.Sprintf("Missing %q parameter!", "text")
				return
			}
			text = texts[0]

			chains, ok := q["chain"]
			if ok {
				chain = chains[0]
			}
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				res.Error = fmt.Sprintf("ParseForm() err: %v", err)
				return
			}
			text = r.FormValue("text")
			chain = r.FormValue("chain")
		}

		if text == "" {
			res.Error = fmt.Sprintf("No input!")
			return
		}
		res.Input = text
		res.Chain = chain

		// detect locale
		locale, err := det.DetectLang(text)
		if err != nil {
			res.Error = err.Error()
			return
		}
		res.Locale = locale

		var trans *translator.Translator
		if chain == "" {
			c := chain
			switch locale {
			case "zh-HK":
				c = "hk2s"
				trans = tm[c]
			case "zh-TW":
				c = "tw2s"
				trans = tm[c]
			case "zh-CN":
				c = "s2hk"
				trans = tm[c]
			}
			res.Chain = c
		} else {
			if !translator.IsValidChain(chain) {
				res.Error = fmt.Sprintf("Invalid chain, valid chains are: %s", strings.Join(translator.ValidChains, ", "))
				return
			}
			trans = tm[chain]
		}

		out, err := trans.Translate(text)
		if err != nil {
			res.Error = err.Error()
			return
		}
		res.Output = out
	}
}
