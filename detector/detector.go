package detector

import (
	"unicode/utf8"

	"github.com/jackytck/jcconv/translator"
)

// NewDetector constructs a new traditional detector.
func NewDetector(size int) (*Detector, error) {
	transST, err := translator.New("s2t")
	if err != nil {
		return nil, err
	}
	transS2HK, err := translator.New("s2hk")
	if err != nil {
		return nil, err
	}
	d := Detector{
		TranslatorST:   transST,
		TranslatorS2HK: transS2HK,
		SampleSize:     size,
	}
	return &d, nil
}

// Detector detects whether the text is simplified or traditional chinese.
type Detector struct {
	TranslatorST   *translator.Translator
	TranslatorS2HK *translator.Translator
	SampleSize     int
}

// DetectLang detects whether the input string is 'zh-CN', 'zh-HK' or 'zh-TW'.
// Default fallback is 'zh-hk'.
func (d *Detector) DetectLang(s string) (string, error) {
	ret := "zh-hk"
	isTrad, err := d.IsTraditional(s)
	if err != nil {
		return ret, err
	}
	if !isTrad {
		return "zh-CN", nil
	}
	isHK, err := d.IsHK(s)
	if err != nil {
		return ret, err
	}
	if isHK {
		return "zh-HK", nil
	}
	return "zh-TW", nil
}

// IsTraditional detects if the string is a traditional chinese.
func (d *Detector) IsTraditional(s string) (bool, error) {
	p, err := d.Detect(s, d.TranslatorST)
	if err != nil {
		return false, err
	}
	if p > 0.97 {
		return true, nil
	}
	return false, nil
}

// IsHK detects if the string is a traditional hong kong chinese.
func (d *Detector) IsHK(s string) (bool, error) {
	p, err := d.Detect(s, d.TranslatorS2HK)
	if err != nil {
		return false, err
	}
	if p > 0.97 {
		return true, nil
	}
	return false, nil
}

// Detect gives the probability of being in traditional chinese if it is a s-to-t translator.
// Or probability of simplified chineses if it is a t-to-s translator.
// Non chinese text would give 1. -1 if error.
func (d *Detector) Detect(s string, t *translator.Translator) (float64, error) {
	// a. translate
	tt, err := t.Translate(s)
	if err != nil {
		return -1, err
	}

	// b. compute sample size
	tsize := utf8.RuneCountInString(tt)
	if d.SampleSize > 0 && tsize > d.SampleSize {
		tsize = d.SampleSize
	}
	if tsize == 0 {
		return -1, nil
	}
	ssize := utf8.RuneCountInString(s)
	if ssize != tsize {
		return 0, nil
	}

	// c. count probability
	rs := []rune(s)
	rt := []rune(tt)
	cnt := 0
	for i, c := range rt {
		if c == rs[i] {
			cnt++
		}
	}

	return float64(cnt) / float64(tsize), nil
}
