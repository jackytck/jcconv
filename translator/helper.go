package translator

import (
	"fmt"

	"github.com/jackytck/jcconv/lib"
)

// ValidChains defines the valid chains.
var ValidChains = []string{"s2hk", "s2tw", "hk2s", "tw2s"}

// IsValidChain tells if the chain key is valid.
func IsValidChain(c string) bool {
	_, valid := contains(ValidChains, c)
	return valid
}

// NewAll creates all types of translators.
func NewAll() (map[string]*Translator, error) {
	// hk
	s2hk, err := New("s2hk")
	if err != nil {
		return nil, err
	}
	hk2s, err := New("hk2s")
	if err != nil {
		return nil, err
	}
	// tw
	s2tw, err := New("s2tw")
	if err != nil {
		return nil, err
	}
	tw2s, err := New("tw2s")
	if err != nil {
		return nil, err
	}
	m := map[string]*Translator{
		"s2hk": s2hk,
		"hk2s": hk2s,
		"s2tw": s2tw,
		"tw2s": tw2s,
	}
	return m, nil
}

// New creates a new translator by name.
// name: "s2hk"
func New(name string) (*Translator, error) {
	switch name {
	case "s2t":
		return newStoT()
	case "s2hk":
		return newStoHK()
	case "s2tw":
		return newStoTW()
	case "hk2s":
		return newHKtoS()
	case "tw2s":
		return newTWtoS()
	default:
		return nil, fmt.Errorf("%q is not valid, try: 's2hk', 's2tw', 'hk2s' or 'tw2s'", name)
	}
}

// newStoT creates a new simplified chinese to traditional Translator
// without any variant.
func newStoT() (*Translator, error) {
	main := []string{
		"STPhrases.txt",
		"STCharacters.txt",
	}
	variant := []string{}
	return newTranslator(main, variant)
}

// newStoHK creates a new simplified chinese to hk Translator.
func newStoHK() (*Translator, error) {
	main := []string{
		"STPhrases.txt",
		"STCharacters.txt",
	}
	variant := []string{
		"HKVariants.txt",
		"HKVariantsPhrases.txt",
		"TWPhrasesIT.txt",
	}
	return newTranslator(main, variant)
}

// newStoTW creates a new simplified chinese to tw Translator.
func newStoTW() (*Translator, error) {
	main := []string{
		"STPhrases.txt",
		"STCharacters.txt",
	}
	variant := []string{
		"TWVariants.txt",
		"TWPhrasesOther.txt",
		"TWPhrasesName.txt",
		"TWPhrasesIT.txt",
	}
	return newTranslator(main, variant)
}

// newHKtoS creates a new hk to simplified chinese Translator.
func newHKtoS() (*Translator, error) {
	main := []string{
		"TWRevPhrasesIT.txt",
		"HKVariantsRevPhrases.txt",
	}
	variant := []string{
		"TSPhrases.txt",
		"TSCharacters.txt",
	}
	return newTranslator(main, variant)
}

// newTWtoS creates a new tw to simplified chinese Translator.
func newTWtoS() (*Translator, error) {
	main := []string{
		"TWRevPhrasesIT.txt",
		"TWRevPhrasesName.txt",
		"TWRevPhrasesOther.txt",
		"TWVariantsRevPhrases.txt",
	}
	variant := []string{
		"TSPhrases.txt",
		"TSCharacters.txt",
	}
	return newTranslator(main, variant)
}

// newTranslator creates a new Translator.
func newTranslator(main, variant []string) (*Translator, error) {
	c, err := lib.NewChain(main, variant)
	if err != nil {
		return nil, err
	}
	t := Translator{c, -1}
	return &t, nil
}

// contains tells whether a contains s.
// And return the index of first match. If not found, index is -1.
func contains(a []string, s string) (int, bool) {
	for i, e := range a {
		if s == e {
			return i, true
		}
	}
	return -1, false
}
