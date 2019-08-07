package lib

import "strings"

// NewChain constructs a new chain from given chains.
// Input: chains is the slice of file names of translations.
func NewChain(chains []string) (*Chain, error) {
	var dicts []*Dict
	for _, c := range chains {
		d, err := NewDict(c)
		if err != nil {
			return nil, err
		}
		dicts = append(dicts, d)
	}
	c := Chain{
		Dicts: dicts,
		Level: len(dicts),
	}
	return &c, nil
}

// Chain represents the chain of dicts of translation
type Chain struct {
	Level int
	Dicts []*Dict
}

// Translate translate the given string according to the chains setup
// NewChain.
func (c *Chain) Translate(input string) (string, error) {
	r := []rune(input)
	var tokens []string
	for i := 0; i < len(r); {
		var token string
		s := r[i:]
		k, v, err := c.match(string(s))
		if err != nil {
			return "", err
		}
		if len([]rune(k)) == 0 {
			token = string(r[i])
			i++
		} else {
			token = v
			i += len([]rune(k))
		}
		tokens = append(tokens, token)
	}
	return strings.Join(tokens, ""), nil
}

// match gets the longest match.
func (c *Chain) match(s string) (string, string, error) {
	max := 0
	maxK, maxM := "", ""
	for _, d := range c.Dicts {
		k, m, err := d.Match(s)
		if err != nil {
			return "", "", err
		}
		// later one override previous one
		if len([]rune(k)) >= max {
			max = len([]rune(k))
			maxK = k
			maxM = m

		}
	}
	return maxK, maxM, nil
}
