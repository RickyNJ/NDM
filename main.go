package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "slices"

    . "github.com/RickyNJ/NDM/mocks"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    m := ReadMappingsDir("mappings/")
    d := GenerateMockDevice(m)

    for {
        fmt.Print(">> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSuffix(input, "\n")
        splitInput := strings.Split(input, " ")
        splitInput = slices.DeleteFunc(splitInput, func(s string) bool {
            return s == ""
        })

        if val, ok := d.Commands[splitInput[0]]; ok {
            output := GetFinalNode(val, splitInput)
            fmt.Print(GetNodeOutput(output))
        } else {
            fmt.Println("command not configured")
        }
    }
}
