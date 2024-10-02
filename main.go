package main

import (
	"fmt"
	"os"
)

type commandNode struct {
	Value  string
	Output interface{}
	Next   []*commandNode
}

func getKeywords(commands []*commandNode) map[string]*commandNode {
	keywordDict := make(map[string]*commandNode)

	for _, command := range commands {
		keywordDict[command.Value] = command
	}

	return keywordDict
}

func getOutput(node *commandNode, args []string) interface{} {
    fmt.Println(node.Value)
    fmt.Println(args)
	if len(args) == 1 && node.Output != nil {
		return node.Output
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i].Value == args[1] {
			getOutput(node.Next[i], args[1:])
		}
	}
	return nil
}

func main() {
	rpdNode := &commandNode{
		Value:  "rpd",
		Output: "showing status rpd",
	}
	ccapNode := &commandNode{
		Value:  "ccap",
		Output: "showing status ccap",
	}
	interfaceNode := &commandNode{
		Value: "interface",
		Next:  []*commandNode{rpdNode, ccapNode},
	}
	showNode := &commandNode{
		Value: "show",
		Next:  []*commandNode{interfaceNode},
	}

	commands := []*commandNode{showNode}

	keywordDict := getKeywords(commands)
	args := os.Args[1:]
    // this should not be needed in the final product
    if len(args) == 0 {
        fmt.Println("no command ran")
        return
    }

	if command, ok := keywordDict[args[0]]; ok {
		fmt.Println(getOutput(command, args))
	}
}
