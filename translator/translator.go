package translator

import (
	"fmt"
	"strings"

	"github.com/jackytck/jcconv/file"
	"github.com/jackytck/jcconv/lib"
)

// Translator represents a particular chain of translation.
type Translator struct {
	Chain  *lib.Chain
	Thread int
}

// TranslateFile translates an input from in and output it to out.
func (t *Translator) TranslateFile(in, out string) error {
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
	s := strings.Join(w, "\n")
	if out == "" {
		fmt.Println(s)
	} else {
		err := file.WriteFile(strings.Join(w, "\n"), out)
		if err != nil {
			return err
		}
	}

	return nil
}

// Translate translates in ram without file io.
// TODO: get error from result channel
func (t *Translator) Translate(line string) (string, error) {
	lines, size := file.SplitLine(line)
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

	s := strings.Join(w, "\n")
	return s, nil
}
