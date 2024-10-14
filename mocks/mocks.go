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

func updateTree(node *mockNode, command []string) {
    if len(command) == 1 {
           //  check if command already has output, if it does not set the output of that node
        return
    }
    updateTree(node, command[1:])
}

func GenerateMockDevice(mocks []*mock) *mockDevice {
    device := &mockDevice{}
    for _, m := range mocks {
        splitCommands := strings.Split(m.Command, " ")
        if _, ok := device.commands[splitCommands[0]]; !ok {
            device.commands[splitCommands[0]] = &mockNode{Value: splitCommands[0]}
        }
        updateTree(device.commands[splitCommands[0]], splitCommands)
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

