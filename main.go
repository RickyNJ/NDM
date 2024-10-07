package main

import (
	"fmt"
	"os"
)

type commandNode struct {
	Value  string
	Next   []*commandNode
	Output interface{}
}

type argumentNode struct {
    commandNode
    ValidAnswers []string
}

func getKeywords(commands []*commandNode) map[string]*commandNode {
	keywordDict := make(map[string]*commandNode)
	for _, command := range commands {
		keywordDict[command.Value] = command
	}

	return keywordDict
}

func getOutput(node *commandNode, args []string) *commandNode {
    fmt.Println(node.Value)
	if len(args) == 1 {
        return node
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i].Value == args[1] {
			return getOutput(node.Next[i], args[1:])
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
        Output: "No arguments given, use show interface --help to see",
	}
	showNode := &commandNode{
		Value: "show",
		Next:  []*commandNode{interfaceNode},
        Output: "No arguments given, use show --help to see",
	}

	commands := []*commandNode{showNode}
	keywordDict := getKeywords(commands)
	args := os.Args[1:]


    if len(args) == 0 {
        fmt.Println("no arguments given")
        return
    }

	if command, ok := keywordDict[args[0]]; ok {
        fmt.Println(command.Value)
		output := getOutput(command, args)
        fmt.Println(output.Output)
	}
}
