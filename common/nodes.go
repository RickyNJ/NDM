package common

import (
)


type Node interface {
    GetOutput(interface{}) string
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

