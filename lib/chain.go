package lib

import "strings"

// NewChain constructs a new chain from given chains.
// Input: main is the slice of file names of the main translations.
// Input: variant is the slice of file names of the variant translations.
func NewChain(main []string, variant []string) (*Chain, error) {
	mc, err := newDicts(main)
	if err != nil {
		return nil, err
	}
	vc, err := newDicts(variant)
	if err != nil {
		return nil, err
	}

	c := Chain{
		Main:    mc,
		Variant: vc,
		Level:   len(mc) + len(vc),
	}
	return &c, nil
}

// Chain represents the chain of dicts of translation
type Chain struct {
	Level   int
	Main    []*Dict
	Variant []*Dict
}

// Translate translate the given string according to the chains setup
// NewChain.
func (c *Chain) Translate(input string) (string, error) {
	out, err := translate(input, c.matchMain)
	if err != nil {
		return "", err
	}
	return translate(out, c.matchVariant)
}

func (c *Chain) matchMain(s string) (string, string, error) {
	return match(s, c.Main)
}

func (c *Chain) matchVariant(s string) (string, string, error) {
	return match(s, c.Variant)
}

func newDicts(chains []string) ([]*Dict, error) {
	var ret []*Dict
	for _, c := range chains {
		d, err := NewDict(c)
		if err != nil {
			return nil, err
		}
		ret = append(ret, d)
	}
	return ret, nil
}

func translate(input string, m func(string) (string, string, error)) (string, error) {
	r := []rune(input)
	var tokens []string
	for i := 0; i < len(r); {
		var token string
		s := r[i:]
		k, v, err := m(string(s))
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
func match(s string, dict []*Dict) (string, string, error) {
	max := 0
	maxK, maxM := "", ""
	for _, d := range dict {
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
