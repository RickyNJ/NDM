package main

import (
	"fmt"

	"github.com/RickyNJ/NDM/mocks"
)
func main() {
    m := mocks.ReadMappingsDir()
    d := mocks.GenerateMockDevice(m)

    for _, n := range d.Commands["show"].Next {
        fmt.Println(n.Value)
    }

    fmt.Println(len(d.Commands))
    return
}
