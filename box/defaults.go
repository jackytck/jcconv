//go:generate go run gen.go

package box

// Has a file in box.
func Has(file string) bool {
	return resources.Has(file)
}

// Get a file from box.
func Get(file string) ([]byte, bool) {
	return resources.Get(file)
}

// Add a file content to box.
func Add(file string, content []byte) {
	resources.Add(file, content)
}

type resourceBox struct {
	storage map[string][]byte
}

// Has tells if the box has the file specified in relative path.
func (r *resourceBox) Has(file string) bool {
	if _, ok := r.storage[file]; ok {
		return true
	}
	return false
}

// Get gets the file content.
// Always use / for looking up
// For example: /init/README.md is actually resources/init/README.md
func (r *resourceBox) Get(file string) ([]byte, bool) {
	if f, ok := r.storage[file]; ok {
		return f, ok
	}
	return nil, false
}

// Add adds a file to the box.
func (r *resourceBox) Add(file string, content []byte) {
	r.storage[file] = content
}

// Resource exposed.
var resources = newResourceBox()

func newResourceBox() *resourceBox {
	return &resourceBox{storage: make(map[string][]byte)}
}
