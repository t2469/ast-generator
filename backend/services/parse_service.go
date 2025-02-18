package services

import (
	"errors"
	sitter "github.com/smacker/go-tree-sitter"
	cppsitter "github.com/smacker/go-tree-sitter/cpp"
	gositter "github.com/smacker/go-tree-sitter/golang"
)

type ASTNode struct {
	Type      string    `json:"type"`
	StartByte int       `json:"start_byte"`
	EndByte   int       `json:"end_byte"`
	Children  []ASTNode `json:"children,omitempty"`
}

func ParseCode(language string, code string) (ASTNode, error) {
	var lang *sitter.Language

	switch language {
	case "go":
		lang = gositter.GetLanguage()
	case "cpp":
		lang = cppsitter.GetLanguage()
	default:
		return ASTNode{}, errors.New("unsupported language")
	}

	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	tree := parser.Parse(nil, []byte(code))
	rootNode := tree.RootNode()

	ast := convertNode(rootNode, []byte(code))
	return ast, nil
}

func convertNode(node *sitter.Node, code []byte) ASTNode {
	astNode := ASTNode{
		Type:      node.Type(),
		StartByte: int(node.StartByte()),
		EndByte:   int(node.EndByte()),
	}

	if node.ChildCount() > 0 {
		children := make([]ASTNode, 0, node.ChildCount())
		for i := 0; i < int(node.ChildCount()); i++ {
			child := node.Child(i)
			if child == nil {
				continue
			}
			children = append(children, convertNode(child, code))
		}
		astNode.Children = children
	}
	return astNode
}
