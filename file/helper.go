package file

import (
	"bufio"
	"os"
)

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
