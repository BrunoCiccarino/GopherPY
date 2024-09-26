package lexer

import (
	"os"
	"strings"
)

func TokenizeFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(data)
	return strings.Fields(content), nil
}
