package file

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

// ScanFile scans the file (located at path) and channels it back per line.
func ScanFile(path string) (<-chan Line, int, <-chan error) {
	lines := make(chan Line)
	errc := make(chan error, 1)

	ss, err := ReadFile(path)
	errc <- err

	go func() {
		defer close(lines)
		for i, s := range ss {
			lines <- Line{int64(i), s}
		}
	}()

	return lines, len(ss), errc
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

// IsTextFile tells if the file is a plain text file.
func IsTextFile(file string) (bool, error) {
	ext, err := GuessFileType(file)
	if err != nil {
		return false, err
	}
	if strings.Contains(ext, "text/") {
		return true, nil
	}
	return false, nil
}

// GuessFileType guesses the type of file.
func GuessFileType(file string) (string, error) {
	buff := make([]byte, 512)
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = f.Read(buff)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buff), nil
}

// IsDir checks if the given path is a directory.
func IsDir(d string) (bool, error) {
	f, err := os.Stat(d)
	if err != nil {
		return false, err
	}
	if f.Mode().IsDir() {
		return true, nil
	}
	return false, nil
}

// EnsureDir ensures a dir exists.
func EnsureDir(p string, perm os.FileMode) error {
	if perm == 0 {
		perm = 0755
	}
	// ensure output dir
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return os.MkdirAll(p, perm)
	}
	return nil
}
