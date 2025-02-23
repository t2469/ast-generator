package services

import (
	"errors"
	"sync/atomic"

	sitter "github.com/smacker/go-tree-sitter"
	bashsitter "github.com/smacker/go-tree-sitter/bash"
	csitter "github.com/smacker/go-tree-sitter/c"
	cppsitter "github.com/smacker/go-tree-sitter/cpp"
	csssitter "github.com/smacker/go-tree-sitter/css"
	dockerfilesitter "github.com/smacker/go-tree-sitter/dockerfile"
	gositter "github.com/smacker/go-tree-sitter/golang"
	htmlsitter "github.com/smacker/go-tree-sitter/html"
	javasitter "github.com/smacker/go-tree-sitter/java"
	jssitter "github.com/smacker/go-tree-sitter/javascript"
	kotlinsitter "github.com/smacker/go-tree-sitter/kotlin"
	phpsitter "github.com/smacker/go-tree-sitter/php"
	pythonsitter "github.com/smacker/go-tree-sitter/python"
	rubysitter "github.com/smacker/go-tree-sitter/ruby"
	rustsitter "github.com/smacker/go-tree-sitter/rust"
	sqlsitter "github.com/smacker/go-tree-sitter/sql"
	yamlsitter "github.com/smacker/go-tree-sitter/yaml"
)

type Point struct {
	Row    uint32 `json:"row"`
	Column uint32 `json:"column"`
}

type ASTNode struct {
	ID         int       `json:"id"`                 // 各ノードに一意なID
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
	case "bash":
		lang = bashsitter.GetLanguage()
	case "c":
		lang = csitter.GetLanguage()
	case "cpp":
		lang = cppsitter.GetLanguage()
	case "css":
		lang = csssitter.GetLanguage()
	case "dockerfile":
		lang = dockerfilesitter.GetLanguage()
	case "go":
		lang = gositter.GetLanguage()
	case "html":
		lang = htmlsitter.GetLanguage()
	case "java":
		lang = javasitter.GetLanguage()
	case "javascript":
		lang = jssitter.GetLanguage()
	case "kotlin":
		lang = kotlinsitter.GetLanguage()
	case "php":
		lang = phpsitter.GetLanguage()
	case "python":
		lang = pythonsitter.GetLanguage()
	case "ruby":
		lang = rubysitter.GetLanguage()
	case "rust":
		lang = rustsitter.GetLanguage()
	case "sql":
		lang = sqlsitter.GetLanguage()
	case "yaml":
		lang = yamlsitter.GetLanguage()
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
