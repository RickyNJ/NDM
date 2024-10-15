package mocks

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type mockDevice struct {
    commands map[string]*mockNode
}

type mockNode struct {
    Value string
    Output string
    Next []*mockNode
}

type mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
}

func updateTree(node *mockNode, command []string, response string) {
    if len(command) == 1 {
        if node.Output != " " {
            panic("this command already has an output set!")
        } else {
            node.Output = response
        }
        return
    }
    for _, n := range node.Next{
        if n.Value == command[1] {
            updateTree(n, command[1:], response)
        }
    }
    nextNode := &mockNode{Value: command[1]}
    node.Next = append(node.Next, nextNode)
    updateTree(nextNode, command[1:], response)
}

func GenerateMockDevice(mocks []*mock) *mockDevice {
    device := &mockDevice{}
    for _, m := range mocks {
        splitCommands := strings.Split(m.Command, " ")
        if _, ok := device.commands[splitCommands[0]]; ok {
            device.commands[splitCommands[0]] = &mockNode{Value: splitCommands[0]}
        }
        updateTree(device.commands[splitCommands[0]], splitCommands, m.Response)
    }
    return device 
}

func GenerateFromJSON(filepath string) *mock {
    m := &mock{} 
    f, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }

    unmarshallingError := json.Unmarshal(f, m)
    if unmarshallingError != nil {
        panic("error while mashaling ")
    }

    return m 
}

func ReadMappingsDir() []*mock {
    var result []*mock
    mappings, err  := os.ReadDir("mappings")
    if err != nil{
        panic("error while opening mappings directory")
    }

    for _, v := range mappings {
        var mapname string
        mapname = "mappings/" + v.Name()
        fmt.Println(mapname)
        m := GenerateFromJSON(mapname)
        result = append(result, m)
    }

    return result
}

