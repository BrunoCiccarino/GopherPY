package generator

import (
	"gopherpy/src/parser"
	"gopherpy/src/transpiler"
	"io/ioutil"
)

func GenerateGoCode(node *parser.Node) string {
	return transpiler.Transpile(node)
}

func SaveGoCode(filename string, code string) error {
	return ioutil.WriteFile(filename, []byte(code), 0644)
}
