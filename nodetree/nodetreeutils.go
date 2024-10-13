package nodetree

import(
    "fmt"
)

type Node interface {
    GetOutput(interface{}) string
}

type Mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
}

type CommandNode struct {
    Value string
    Next []*CommandNode
    Output string
    OutputPath string
}

type argumentNode struct {
    CommandNode
    ValidAnswers []string
}

func ReadFile(path string) string {
    return ""
}

func (n *CommandNode) GetOutput() string {
    if n.Output !=  ""{
        return n.Output
    }
    if n.OutputPath != "" {
        return ReadFile(n.OutputPath)
    }
    return "" 
}
func GetKeywords(commands []*CommandNode) map[string]*CommandNode {
	keywordDict := make(map[string]*CommandNode)
	for _, command := range commands {
		keywordDict[command.Value] = command
	}

	return keywordDict
}

func GetFinalNode(node *CommandNode, args []string) *CommandNode {
    fmt.Println(node.Value)
	if len(args) == 1 {
        return node
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i].Value == args[1] {
			return GetFinalNode(node.Next[i], args[1:])
		}
	}
	return nil
}

// This function loads the json file, and outputs the array of rootnodes
// Add check so that node cannot be generated with both a output and a outputfile
// func LoadTrees(filepath string) []*CommandNode {
//
// }
