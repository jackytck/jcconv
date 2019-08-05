package lib

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
func (c *Chain) Translate(s string) (string, error) {
	// @TODO
	return "", nil
}
