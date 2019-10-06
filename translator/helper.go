package translator

import (
	"fmt"

	"github.com/jackytck/jcconv/lib"
)

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
