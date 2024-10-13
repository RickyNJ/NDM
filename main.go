package main

import (
	"fmt"

	"github.com/RickyNJ/NDM/mocks"
)
func main() {
    mocks := mocks.ReadMappingsDir()
    for _, m := range mocks {
        fmt.Println(m.Response) 
    }
    return
}
// func main() {
// 	rpdNode := &CommandNode{
// 		Value:  "rpd",
// 		Output: "showing status rpd",
// 	}
// 	ccapNode := &CommandNode{
// 		Value:  "ccap",
// 		Output: "showing status ccap",
// 	}
// 	interfaceNode := &CommandNode{
// 		Value: "interface",
// 		Next:  []*CommandNode{rpdNode, ccapNode},
//         Output: "No arguments given, use show interface --help to see",
// 	}
// 	showNode := &CommandNode{
// 		Value: "show",
// 		Next:  []*CommandNode{interfaceNode},
//         Output: "No arguments given, use show --help to see",
// 	}
//
// 	commands := []*CommandNode{showNode}
// 	keywordDict := GetKeywords(commands)
// 	args := os.Args[1:]
//
//
//     if len(args) == 0 {
//         fmt.Println("no arguments given")
//         return
//     }
//
// 	if command, ok := keywordDict[args[0]]; ok {
//         fmt.Println(command.Value)
// 		output := getOutput(command, args)
//         fmt.Println(output.Output)
// 	}
// }
