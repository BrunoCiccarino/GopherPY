package lexer

import (
	"io/ioutil"
	"strings"
)

func TokenizeFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(data)
	return strings.Fields(content), nil
}
