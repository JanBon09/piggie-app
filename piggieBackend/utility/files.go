package utility

import (
	"bufio"
	"io"
	"os"
)

// Reading file with path specified under filepath variable
// Returning file content in form of one string
func ReadFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	w := bufio.NewReader(file)

	encodedSecretKey, err := w.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			return "", err
		}
	}

	return encodedSecretKey, nil
}
