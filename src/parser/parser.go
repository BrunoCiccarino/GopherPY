package parser

import (
	"strings"
)

type Node struct {
	Type     string
	Value    string
	Children []*Node
}

func Parse(tokens []string) *Node {
	if len(tokens) == 0 {
		return nil
	}

	if tokens[0] == "def" {
		return parseFunction(tokens)
	}

	if strings.HasPrefix(tokens[0], "if") {
		return parseIfStatement(tokens)
	}

	return nil
}

func parseFunction(tokens []string) *Node {
	name := tokens[1]
	body := tokens[4:]
	return &Node{
		Type:  "Function",
		Value: name,
		Children: []*Node{
			{
				Type:  "Body",
				Value: strings.Join(body, " "),
			},
		},
	}
}

func parseIfStatement(tokens []string) *Node {
	return &Node{
		Type:  "IfStatement",
		Value: strings.Join(tokens, " "),
	}
}
