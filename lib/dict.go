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

// Match matches the longest prefix match in the trie.
// Return the key and value matched.
func (d *Dict) Match(s string) (string, string, error) {
	mats, err := d.MatchAll(s)
	if err != nil {
		return "", "", err
	}
	maxLen := 0
	maxKey := ""
	maxVal := ""
	for k, v := range mats {
		if len(k) > maxLen {
			maxLen = len(k)
			maxKey = k
			maxVal = v[0] // only use the first variant
		}
	}

	return maxKey, maxVal, nil
}

// MatchAll prefix matches all match via its trie from the given s.
func (d *Dict) MatchAll(s string) (map[string][]string, error) {
	if d.trie == nil {
		return nil, errors.New("invalid dict")
	}
	ret := make(map[string][]string)
	// If `num` is 0, it returns all matches.
	num := 0
	for _, id := range d.trie.PrefixMatch([]byte(s), num) {
		k, err := d.trie.Key(id)
		if err != nil {
			return nil, err
		}
		v, err := d.trie.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(k)] = d.translations[v]
	}
	return ret, nil
}
