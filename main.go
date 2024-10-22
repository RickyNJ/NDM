package main

import (
    // "bufio"
    "fmt"
    // "os"
    //    "strings"

    . "github.com/RickyNJ/NDM/mocks"
)

func main() {
    // reader := bufio.NewReader(os.Stdin)

    m := ReadMappingsDir("mappings/")
    d := GenerateMockDevice(m)
    input := []string{"show", "interface"}
    final := GetFinalNode(d.Commands["show"], input)
    fmt.Println(final.Output)
    // for {
    //     fmt.Print(">> ")
    //     input, _ := reader.ReadString('\n')
    //     splitInput := strings.Split(input, " ")
    //     fmt.Println(len(splitInput))
    //
    //     if val, ok := d.Commands[splitInput[0]]; ok {
    //         output := GetFinalNode(val, splitInput)
    //         fmt.Println(output.Output)
    //     } else {
    //         fmt.Println("command not configured")
    //     }
    // }
}
