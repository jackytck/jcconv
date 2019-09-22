package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackytck/jcconv/detector"
	"github.com/jackytck/jcconv/translator"
)

// Translate auto translates the input text.
func Translate(det *detector.Detector, trans2hk, trans2s *translator.Translator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		var res TransRes
		defer func() {
			res.Elapsed = time.Since(start).String()
			json.NewEncoder(w).Encode(res)
		}()

		var text string
		switch r.Method {
		case http.MethodGet:
			texts, ok := r.URL.Query()["text"]
			if !ok || len(texts[0]) < 1 {
				res.Error = fmt.Sprintf("Missing %q parameter!", "text")
				return
			}
			text = texts[0]
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				res.Error = fmt.Sprintf("ParseForm() err: %v", err)
				return
			}
			text = r.FormValue("text")
		}

		if text == "" {
			res.Error = fmt.Sprintf("No input!")
			return
		}
		res.Input = text

		isTrad, err := det.IsTraditional(text)
		if err != nil {
			res.Error = err.Error()
			return
		}
		var trans *translator.Translator
		if isTrad {
			trans = trans2s
		} else {
			trans = trans2hk
		}

		out, err := trans.TranslateOne(text)
		if err != nil {
			res.Error = err.Error()
			return
		}
		res.Output = out
	}
}
