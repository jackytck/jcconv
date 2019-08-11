package file

import (
	"bufio"
	"os"
)

// ScanFile scans the file (located at path) and channels it back per line.
func ScanFile(path string) (<-chan Line, <-chan error) {
	lines := make(chan Line)
	errc := make(chan error, 1)

	go func() {
		defer close(lines)

		ss, err := ReadFile(path)
		errc <- err

		for i, s := range ss {
			lines <- Line{int64(i), s}
		}
	}()

	return lines, errc
}

// ReadFile reads a text file from path.
func ReadFile(path string) ([]string, error) {
	var ret []string
	file, err := os.Open(path)
	if err != nil {
		return ret, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return ret, err
	}

	return ret, nil
}

// WriteFile writes text to file at path.
func WriteFile(text, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}
