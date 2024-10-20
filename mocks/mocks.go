package mocks

import (
    "fmt"
	"encoding/json"
	"os"
	"strings"
    "log"
)

type MockDevice struct {
    Commands map[string]*MockNode
}

type MockNode struct {
    Value string
    Output string
    Next []*MockNode
}

type Mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
}

func updateTree(node *MockNode, command []string, response string) {
    log.Println("TREEUPDATE: tree with node: ", node.Value, "command: ", command)
    if len(command) == 1 {
        log.Println("Final element of command")
        if node.Output != "" {
            panic("this command already has an output set!")
        } else {
            log.Println("Set output: ",response)
            node.Output = response
        }
        return
    }
    var nextNode *MockNode
    for _, n := range node.Next{
        if n.Value == command[1] {
            log.Println("Found node in next with same value")
            nextNode = n
        }
    }

    if nextNode == nil {
        nextNode = &MockNode{ Value: command[1]}
        node.Next = append(node.Next, nextNode)
    }

    updateTree(nextNode, command[1:], response)
}

func GenerateMockDevice(mocks []*Mock) *MockDevice {
    device := &MockDevice{Commands: make(map[string]*MockNode)}

    for _, m := range mocks {
        log.Println("NEWMOCK: tree with mock: ", m.Command)
        splitCommands := strings.Split(m.Command, " ")
        if _, ok := device.Commands[splitCommands[0]]; !ok {
            log.Println("NEWROOTCOMMAND: ", splitCommands[0], " Generating a node" )
            device.Commands[splitCommands[0]] = &MockNode{Value: splitCommands[0]}
        } else {
            log.Println("USEEXISTINGROOT")
        }
        updateTree(device.Commands[splitCommands[0]], splitCommands, m.Response)
    }

    return device 
}

func GenerateFromJSON(filepath string) *Mock {
    m := &Mock{} 
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

func ReadMappingsDir(dir string) []*Mock {
    var result []*Mock
    mappings, err  := os.ReadDir(dir)
    if err != nil{
        panic("error while opening mappings directory")
    }

    for _, v := range mappings {
        var mapname string
        mapname = dir + v.Name()
        log.Println("found:", mapname)
        m := GenerateFromJSON(mapname)
        result = append(result, m)
    }

    return result
}

func GetFinalNode(node *MockNode, args []string) *MockNode {
    fmt.Println(node.Value)
	if len(args) == 1 {
        return node
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i].Value == args[1] {
			return GetFinalNode(node.Next[i], args[1:])
		}
	}

	return nil
}
