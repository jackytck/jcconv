package lib

import (
	"errors"
	"strings"

	"github.com/jackytck/go-chinese-converter/box"
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
		return key, trans, errors.New("translation not found")
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
