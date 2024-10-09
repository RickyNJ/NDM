package common

type CommandNode struct {
    Value string
    Next []*CommandNode
    Output interface{}
}

type argumentNode struct {
    CommandNode
    ValidAnswers []string
}
