package envload

import (
	"bufio"
	"bytes"
	"os"
)

func File(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		// Ignore comments and empty lines
		trimmed := bytes.TrimSpace(line)
		if len(trimmed) == 0 || bytes.HasPrefix(trimmed, []byte("#")) {
			continue
		}

		// Split the line into key and value
		parts := bytes.SplitN(trimmed, []byte("="), 2)
		if len(parts) != 2 {
			continue
		}

		key := bytes.TrimSpace(parts[0])
		value := bytes.TrimSpace(parts[1])

		// Remove optional surrounding quotes
		value = bytes.Trim(value, `"'`)

		// Set the environment variable
		os.Setenv(string(key), string(value))
	}

	return scanner.Err()
}
