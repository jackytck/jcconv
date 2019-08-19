package translator

import (
	"fmt"

	"github.com/jackytck/go-chinese-converter/lib"
)

// New creates a new translator by name.
// name: "s2hk"
func New(name string) (*Translator, error) {
	switch name {
	case "s2hk":
		return newStoHK()
	case "s2tw":
		return newStoTW()
	default:
		return nil, fmt.Errorf("%q not found", name)
	}

}

// NewStoHK creates a new simplified chinese to hk Translator.
func newStoHK() (*Translator, error) {
	main := []string{
		"STPhrases.txt",
		"STCharacters.txt",
	}
	variant := []string{
		"HKVariants.txt",
		"HKVariantsPhrases.txt",
		"HKVariantsRevPhrases.txt",
		"TWPhrasesIT.txt",
	}
	return newTranslator(main, variant)
}

// NewStoTW creates a new simplified chinese to tw Translator.
func newStoTW() (*Translator, error) {
	main := []string{
		"STPhrases.txt",
		"STCharacters.txt",
	}
	variant := []string{
		"TWVariants.txt",
		"TWVariantsRevPhrases.txt",
		"TWPhrasesOther.txt",
		"TWPhrasesName.txt",
		"TWPhrasesIT.txt",
	}
	return newTranslator(main, variant)
}

// NewTranslator creates a new Translator.
func newTranslator(main, variant []string) (*Translator, error) {
	c, err := lib.NewChain(main, variant)
	if err != nil {
		return nil, err
	}
	t := Translator{c, -1}
	return &t, nil
}
