package services

import (
	"errors"
	"sync/atomic"

	sitter "github.com/smacker/go-tree-sitter"
	cppsitter "github.com/smacker/go-tree-sitter/cpp"
	gositter "github.com/smacker/go-tree-sitter/golang"
)

type Point struct {
	Row    uint32 `json:"row"`
	Column uint32 `json:"column"`
}

type ASTNode struct {
	ID         int       `json:"id"`                 // 各ノードに一意なID（
	Type       string    `json:"type"`               // ノードの種類（例："identifier", "function_declaration" など）
	Content    string    `json:"content,omitempty"`  // ソースコード上の該当部分のテキスト内容
	StartByte  int       `json:"start_byte"`         // ノードが開始するバイト位置（ソースコード内）
	EndByte    int       `json:"end_byte"`           // ノードが終了するバイト位置（ソースコード内）
	StartPoint Point     `json:"start_point"`        // ノードの開始位置（行番号と列番号）
	EndPoint   Point     `json:"end_point"`          // ノードの終了位置（行番号と列番号）
	IsExtra    bool      `json:"is_extra"`           // 補助的/余分なノードの場合にtrue（文法的に不要な部分など）
	Children   []ASTNode `json:"children,omitempty"` // 子ノードのリスト（再帰的なツリー構造）
}

var globalNodeID int32 = 0

func nextNodeID() int {
	return int(atomic.AddInt32(&globalNodeID, 1))
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
		ID:         nextNodeID(),
		Type:       node.Type(),
		Content:    string(code[node.StartByte():node.EndByte()]),
		StartByte:  int(node.StartByte()),
		EndByte:    int(node.EndByte()),
		StartPoint: Point{Row: node.StartPoint().Row, Column: node.StartPoint().Column},
		EndPoint:   Point{Row: node.EndPoint().Row, Column: node.EndPoint().Column},
		IsExtra:    node.IsExtra(),
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
