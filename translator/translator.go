package translator

import (
	"strings"

	"github.com/jackytck/go-chinese-converter/file"
	"github.com/jackytck/go-chinese-converter/lib"
)

// NewStoHK creates a new simplified chinese to hk Translator.
func NewStoHK() (*Translator, error) {
	main := []string{
		"STPhrases.txt",
		"STCharacters.txt",
	}
	variant := []string{
		"HKVariants.txt",
		"HKVariantsPhrases.txt",
		"HKVariantsRevPhrases.txt",
	}
	return NewTranslator(main, variant)
}

// NewTranslator creates a new Translator.
func NewTranslator(main, variant []string) (*Translator, error) {
	c, err := lib.NewChain(main, variant)
	if err != nil {
		return nil, err
	}
	t := Translator{c, -1}
	return &t, nil
}

// Translator represents a particular chain of translation.
type Translator struct {
	Chain  *lib.Chain
	Thread int
}

// Translate translates an input from in and output it to out.
func (t *Translator) Translate(in, out string) error {
	lines, size, errc := file.ScanFile(in)
	done := make(chan struct{})
	defer close(done)
	result := make(chan file.Line)

	digester := file.Digester{
		Chain:  t.Chain,
		Done:   done,
		Lines:  lines,
		Result: result,
	}
	digester.Run(t.Thread)

	w := make([]string, size)
	for l := range result {
		w[l.LineNum] = l.Text
	}
	if err := <-errc; err != nil {
		return err
	}
	err := file.WriteFile(strings.Join(w, "\n"), out)
	if err != nil {
		return err
	}

	return nil
}
