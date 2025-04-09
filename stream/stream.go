package stream

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ProcessStream(reader io.Reader, lineHandler func(string)) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lineHandler(line)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}
	return nil
}
