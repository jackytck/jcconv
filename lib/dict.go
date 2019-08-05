package lib

import (
	"errors"

	"github.com/go-ego/cedar"
)

// NewDict constructs a new dictionary from given translation file.
// Input: trans is the file name of translation without '/'.
func NewDict(trans string) (*Dict, error) {
	key, val, err := ReadTrans(trans)
	if err != nil {
		return nil, err
	}
	if len(key) != len(val) {
		return nil, errors.New("key and values mismatched")
	}

	trie := cedar.New()
	for i, k := range key {
		err := trie.Insert([]byte(k), i)
		if err != nil {
			return nil, err
		}
	}

	d := Dict{
		trie:         trie,
		translations: val,
	}
	return &d, nil
}

// Dict represents the dictionary for translating words and phrases.
type Dict struct {
	trie         *cedar.Cedar
	translations [][]string
}

// Match prefix matches the given s via its trie.
func (d *Dict) Match(s string) (map[string][]string, error) {
	// @TODO
	ret := make(map[string][]string)
	return ret, nil
}
