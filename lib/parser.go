package lib

import (
	"fmt"
	"strings"

	"github.com/jackytck/jcconv/box"
)

// ReadTrans reads the translation files.
// Return the key and corresponding 1-to-many translations and error.
func ReadTrans(t string) ([]string, [][]string, error) {
	// var err error
	var key []string
	var trans [][]string

	// a. read bytes
	tv, ok := box.Get("/" + t)
	if !ok {
		return key, trans, fmt.Errorf("%q translation not found", t)
	}

	// b. parse line by line
	for _, line := range strings.Split(string(tv), "\n") {
		toks := strings.Split(strings.TrimSpace(line), "\t")
		if len(toks) < 2 {
			continue
		}
		key = append(key, toks[0])
		trans = append(trans, strings.Fields(toks[1]))
	}

	return key, trans, nil
}
