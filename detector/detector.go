package detector

import (
	"unicode/utf8"

	"github.com/jackytck/jcconv/translator"
)

// NewDetector constructs a new traditional detector.
func NewDetector(size int) (*Detector, error) {
	chain := "s2hk"
	trans, err := translator.New(chain)
	if err != nil {
		return nil, err
	}
	d := Detector{
		Translator: trans,
		SampleSize: size,
	}
	return &d, nil
}

// Detector detects whether the text is simplified or traditional chinese.
type Detector struct {
	Translator *translator.Translator
	SampleSize int
}

// IsTraditional detects if the string is a traditional chinese.
func (d *Detector) IsTraditional(s string) (bool, error) {
	p, err := d.Detect(s)
	if err != nil {
		return false, err
	}
	if p > 0.97 {
		return true, nil
	}
	return false, nil
}

// Detect gives the probability of being in traditional/simplified chinese (according to translator) if it is chinese.
// Non chinese text would give 1. -1 if error.
func (d *Detector) Detect(s string) (float64, error) {
	// a. translate
	t, err := d.Translator.TranslateOne(s)
	if err != nil {
		return -1, err
	}

	// b. compute sample size
	size := utf8.RuneCountInString(t)
	if d.SampleSize > 0 && size > d.SampleSize {
		size = d.SampleSize
	}
	if size == 0 {
		return -1, nil
	}

	// c. count probability
	rs := []rune(s)
	rt := []rune(t)
	cnt := 0
	for i, c := range rt {
		if c == rs[i] {
			cnt++
		}
	}

	return float64(cnt) / float64(size), nil
}
