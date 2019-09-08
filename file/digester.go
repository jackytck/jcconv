package file

import (
	"runtime"
	"sync"

	"github.com/jackytck/jcconv/lib"
)

// Digester digests a single text file with n goroutines.
type Digester struct {
	Chain  *lib.Chain
	Done   <-chan struct{}
	Lines  <-chan Line
	Result chan<- Line
}

// Run starts n number of goroutines to traslate the lines.
// If n is not positive, it will be set to number of CPU cores x 4.
// Return n.
func (d *Digester) Run(n int) int {
	if n <= 0 {
		n = runtime.NumCPU() * 4
	}
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			d.digest()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(d.Result)
	}()

	return n
}

// digest digests all the input lines, translate them and gives back the
// translated result.
func (d *Digester) digest() {
	for line := range d.Lines {
		select {
		case d.Result <- d.translate(line):
		case <-d.Done:
			return
		}
	}
}

// translate translates a line.
func (d *Digester) translate(line Line) Line {
	out, err := d.Chain.Translate(line.Text)
	if err != nil {
		panic(err)
	}
	return Line{line.LineNum, out}
}
