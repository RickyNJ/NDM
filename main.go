package main

import (
	"fmt"

	"github.com/RickyNJ/NDM/mocks"
)
func main() {
    mocks := mocks.ReadMappingsDir()
    for _, m := range mocks {
        fmt.Println(m.Response) 
    }
    return
}
