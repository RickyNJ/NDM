package main

import (
	"fmt"
	"os"

	. "github.com/RickyNJ/NDM/common"
)


func getKeywords(commands []*CommandNode) map[string]*CommandNode {
	keywordDict := make(map[string]*CommandNode)
	for _, command := range commands {
		keywordDict[command.Value] = command
	}

	return keywordDict
}

func getOutput(node *CommandNode, args []string) *CommandNode {
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
	rpdNode := &CommandNode{
		Value:  "rpd",
		Output: "showing status rpd",
	}
	ccapNode := &CommandNode{
		Value:  "ccap",
		Output: "showing status ccap",
	}
	interfaceNode := &CommandNode{
		Value: "interface",
		Next:  []*CommandNode{rpdNode, ccapNode},
        Output: "No arguments given, use show interface --help to see",
	}
	showNode := &CommandNode{
		Value: "show",
		Next:  []*CommandNode{interfaceNode},
        Output: "No arguments given, use show --help to see",
	}

	commands := []*CommandNode{showNode}
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
