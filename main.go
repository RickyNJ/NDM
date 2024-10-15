package main

import (
	"fmt"

	"github.com/RickyNJ/NDM/mocks"
)
func main() {
    m := mocks.ReadMappingsDir()
    d := mocks.GenerateMockDevice(m)

    fmt.Println(d.Commands["show"])

    for _, n := range d.Commands["show"].Next {
        fmt.Println(n.Value)
    }
    return
}
