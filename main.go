package main

import (
	"container/list"
	"fmt"

	"github.com/RickyNJ/NDM/mocks"
)
func ReadTree(root *mocks.MockNode)[]string{
    q := list.New()
    result := []string{}

    q.PushBack(root)

    for q.Len() > 0 {
        current := q.Front().Value.(*mocks.MockNode)
        result = append(result, current.Value)
        q.Remove(q.Front())
        for _, m := range current.Next {
            q.PushBack(m)
        }
    }
    return result
} 

func GetCommandTree(device *mocks.MockDevice, command string) []string {
    return ReadTree(device.Commands["command"])
}

func main() {
    m := mocks.ReadMappingsDir()
    d := mocks.GenerateMockDevice(m)


    fmt.Println(len(d.Commands))
    fmt.Println(ReadTree(d.Commands["show"]))
    return
}
