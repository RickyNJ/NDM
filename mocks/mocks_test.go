package mocks

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func ReadTree(root *MockNode)[]string{
    q := list.New()
    result := []string{}

    q.PushBack(root)

    for q.Len() > 0 {
        current := q.Front().Value.(*MockNode)
        result = append(result, current.Value)
        q.Remove(q.Front())
        for _, m := range current.Next {
            q.PushBack(m)
        }
    }
    return result
} 

func GetCommandTree(device *MockDevice, command string) []string {
    return ReadTree(device.Commands["command"])
}

func TestTreeCreation(t *testing.T) {
    m := ReadMappingsDir("../mappings/")
    d := GenerateMockDevice(m)

    fmt.Println(len(d.Commands))
    showtree := ReadTree(d.Commands["show"])

    if !reflect.DeepEqual(showtree, []string{"show", "interface", "cable", "kj"}) {
        t.Fatal("not equal")
    }
}
