package mocks

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

type MockDevice struct {
    Commands map[string]*MockNode
}

type ArgNode struct {
    Value string
    ValidAnswers []string
    Output string
    OutputFile string
    Next []*MockNode
    Args []*ArgNode
}

type MockNode struct {
    Value string
    Output string
    OutputFile string
    Next []*MockNode
    Args []*ArgNode
}

type Mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
    ResponseFile string `json:"responsefile"`
}



func GetResponse(device *MockDevice, command []string) string {
    if val, ok := device.Commands[command[0]]; ok {
        output := GetFinalNode(val, command)
        return GetNodeOutput(output)
    } else {
        return "command not configured"
    }
}

func checkForMatchingNode(node *MockNode, nextArg string) *MockNode {
    for _, n := range node.Next{
        if n.Value == nextArg {
            log.Println("Found node in next with same value")
            return n
        }
    }
    return nil
}

func checkForMatchingArgument(node *MockNode, nextArg string) *ArgNode {
    for _, n := range node.Args{
          if slices.Contains(n.ValidAnswers, nextArg) {
              return n
          }
    }
    return nil
}

func updateTree(node *MockNode, command []string, response string, responsefile string) {
    log.Println("TREEUPDATE: tree with node: ", node.Value, "command: ", command)
    if len(command) == 1 {
        log.Println("Final element of command")
        if node.Output != "" {
            panic("this command already has an output set!")
        } else {
            log.Println("Set output: ",response)
            node.Output = response
            node.OutputFile = responsefile
        }
        return
    }

    var nextNode *MockNode
    nextNode = checkForMatchingNode(node, command[1])

    if nextNode == nil {
        
    }

    if nextNode == nil {
        nextNode = &MockNode{ Value: command[1]}
        node.Next = append(node.Next, nextNode)
    }

    updateTree(nextNode, command[1:], response, responsefile)
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
        updateTree(device.Commands[splitCommands[0]], splitCommands, m.Response, m.ResponseFile)
    }

    return device 
}

func generateFromJSON(filepath string) *Mock {
    m := &Mock{} 
    f, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }
    
    unmarshallingError := json.Unmarshal(f, m)
    if unmarshallingError != nil {
        panic("error while mashaling ")
    }
    log.Println("Generated mock from json: ", filepath, "\nCommand: ", m.Command, " \nResponse: ", m.Response, "\nResponseFile: ", m.ResponseFile )

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
        m := generateFromJSON(mapname)
        result = append(result, m)
    }

    return result
}

func GetFinalNode(node *MockNode, args []string) *MockNode {
	if len(args) == 1 {
        return node
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i].Value == args[1] {
			return GetFinalNode(node.Next[i], args[1:])
		}
	}

	return &MockNode{Output: "this command has not been configured"}
}

func readOutputfile(node *MockNode) string {
    filepath := "__files/" + node.OutputFile
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatalf("could not open response file ", node.OutputFile, err )
    }
    defer file.Close()

    response, err := io.ReadAll(file) 
    if err != nil {
        log.Fatalf("could not read content of response file ", node.OutputFile, err)
    }

    return string(response)

}

func GetNodeOutput(node *MockNode) string {
    if node.Output != "" {
        return node.Output + "\n"
    }

    if node.OutputFile != "" {
        return readOutputfile(node)
    } 

    return "this command has no output configured"
}

