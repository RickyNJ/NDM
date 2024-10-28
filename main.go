package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "slices"

    "github.com/RickyNJ/NDM/mocks"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    m := mocks.ReadMappingsDir("__mappings/")
    d := mocks.GenerateMockDevice(m)

    for {
        fmt.Print(">> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSuffix(input, "\n")
        
        splitInput := strings.Split(input, " ")
        splitInput = slices.DeleteFunc(splitInput, func(s string) bool {
            return s == ""
        })

        if len(splitInput) > 0 {
            fmt.Print(mocks.GetResponse(d, splitInput))
        }
    }
}
