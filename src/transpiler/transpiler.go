package transpiler

import (
	"fmt"
	"gopherpy/src/parser"
)

func Transpile(node *parser.Node) string {
	if node == nil {
		return ""
	}

	switch node.Type {
	case "Function":
		return transpileFunction(node)
	case "IfStatement":
		return transpileIf(node)
	default:
		return ""
	}
}

func transpileFunction(node *parser.Node) string {
	return fmt.Sprintf("func %s() {\n\t%s\n}", node.Value, node.Children[0].Value)
}

func transpileIf(node *parser.Node) string {
	return fmt.Sprintf("if %s {\n}", node.Value)
}
