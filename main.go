package main

import (
	"fmt"
	"gopherpy/src/generator"
	"gopherpy/src/lexer"
	"gopherpy/src/parser"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <python-file>")
		return
	}

	pythonFilename := os.Args[1]
	goFilename := strings.Replace(pythonFilename, ".py", ".go", 1)

	tokens, err := lexer.TokenizeFile(pythonFilename)
	if err != nil {
		fmt.Println("Error reading Python file:", err)
		return
	}

	ast := parser.Parse(tokens)
	goCode := generator.GenerateGoCode(ast)

	err = generator.SaveGoCode(goFilename, goCode)
	if err != nil {
		fmt.Println("Error saving Go file:", err)
		return
	}

	fmt.Printf("Go code saved to %s\n", goFilename)
}
