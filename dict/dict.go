package dict

import (
	"log"

	"github.com/go-ego/cedar"
)

// NewDict constructs a new dictionary from given chains.
func NewDict(chains []string) (*Dict, error) {
	for _, c := range chains {
		k, v, err := ReadTrans(c)
		if err != nil {
			return nil, err
		}
		// @TODO
		log.Println("key", k)
		log.Println("value", v)
	}
	return nil, nil
}

// Dict represents the dictionary for translating words and phrases.
type Dict struct {
	chains       []string
	trie         *cedar.Cedar
	translations [][]string
}
