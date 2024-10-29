package mocks

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

)

type MockDevice struct {
    Commands map[string]*BaseNode
}

type BaseNode struct {
    Value string
    Output string
    OutputFile string
    Next []*BaseNode
    VariableNext []*VarNode
}

type VarNode struct {
    *BaseNode
    ValidAnswers []string
}

type Mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
    ResponseFile string `json:"responsefile"`
}

type Node interface {
    CheckIfValid(string) bool
}

func Check(n Node, input string) bool {
    return n.CheckIfValid(input)
}


func (n *BaseNode) CheckIfValid(input string) bool {
    if input == n.Value {
        return true
    }
    return false
}

func GetResponse(device *MockDevice, command []string) string {
    if val, ok := device.Commands[command[0]]; ok {
        output := GetFinalNode(val, command)
        return GetNodeOutput(output)
    } else {
        return "command not configured\n"
    }
}

func matchVarNode(node *BaseNode, nextArg string) *VarNode {
    var next *VarNode

    for _, n := range node.VariableNext {
        if n.Value == nextArg {
            log.Println("Found node in next with same value")
            next = n
        }
    }
    if next == nil {
        // create default node here 
        nextBase := &BaseNode{Value: nextArg}
        next = &VarNode{BaseNode: nextBase}
        node.VariableNext = append(node.VariableNext, next)
    }
    return next
}

func matchDefaultNode(node *BaseNode, nextArg string) *BaseNode {
    log.Println("matching default node: ", node.Value)
    log.Println("argument value to match: ", nextArg)
    var next *BaseNode

    for _, n := range node.Next {
        if n.Value == nextArg {
            next = n
        }
    }

    if next == nil {
        next = &BaseNode{Value: nextArg}
    }
    return next
}


func updateTree(node *BaseNode, command []string, response string, responsefile string) {
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

    var nextNode *BaseNode
    runes := []rune(command[1])
    switch runes[0] {
    case '{':
        if runes[len(runes)-1] != '}' {
            panic("either unclosed argument bracket \"{}\" or no space after variable declaration in mappings file")
        }
        nextNode = matchVarNode(node, command[1]).BaseNode
    case '-':
        //  Implement later
    default:
        nextNode = matchDefaultNode(node, command[1])
    }

    updateTree(nextNode, command[1:], response, responsefile)
}

func GenerateMockDevice(mocks []*Mock) *MockDevice {
    device := &MockDevice{Commands: make(map[string]*BaseNode)}

    for _, m := range mocks {
        log.Println("NEWMOCK: tree with mock: ", m.Command)
        splitCommands := strings.Split(m.Command, " ")
        if _, ok := device.Commands[splitCommands[0]]; !ok {
            log.Println("NEWROOTCOMMAND: ", splitCommands[0], " Generating a node" )
            device.Commands[splitCommands[0]] = &BaseNode{Value: splitCommands[0]}
        } else {
            log.Println("USEEXISTINGROOT")
        }
        updateTree(device.Commands[splitCommands[0]], splitCommands, m.Response, m.ResponseFile)
    }

    return device 
}

func generateFromJSON(filepath string) *Mock {
    // add first element of a command is not allowed to be a variable or a flag
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
    mappings, err := os.ReadDir(dir)
    if err != nil{
        panic(err)
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

func GetFinalNode(node *BaseNode, args []string) *BaseNode {
	if len(args) == 1 {
        return node
	}

	for i := 0; i < len(node.Next); i++ {
		if node.Next[i].Value == args[1] {
			return GetFinalNode(node.Next[i], args[1:])
		}
	}

	return &BaseNode{Output: "this command has not been configured"}
}

func readOutputfile(node *BaseNode) string {
    filepath := "__files/" + node.OutputFile
    file, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    response, err := io.ReadAll(file) 
    if err != nil {
        panic(err)
    }

    return string(response)

}

func GetNodeOutput(node *BaseNode) string {
    if node.Output != "" {
        return node.Output + "\n"
    }

    if node.OutputFile != "" {
        return readOutputfile(node)
    } 

    return "this command has no output configured"
}

