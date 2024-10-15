package mocks

import (
	"encoding/json"
	"os"
	"strings"
    "log"
)

type mockDevice struct {
    Commands map[string]*mockNode
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
    log.Println("updating tree with node: ", node.Value, "command: ", command)
    if len(command) == 1 {
        if node.Output != "" {
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
    device := &mockDevice{Commands: make(map[string]*mockNode)}
    for _, m := range mocks {
        log.Println("Updating tree with mock: ", m.Command)
        splitCommands := strings.Split(m.Command, " ")
        if _, ok := device.Commands[splitCommands[0]]; !ok {
            log.Println("No root node found for: ", splitCommands[0], " Generating a node")
            device.Commands[splitCommands[0]] = &mockNode{Value: splitCommands[0]}
        }
        log.Println("Rootnode found:")
        updateTree(device.Commands[splitCommands[0]], splitCommands, m.Response)
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
    log.Println("Generated mock from json: ", filepath, "\nCommand: ", m.Command, " \nResponse: ", m.Response)
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
        log.Println("found:", mapname)
        m := GenerateFromJSON(mapname)
        result = append(result, m)
    }

    return result
}

