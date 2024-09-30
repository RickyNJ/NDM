package main

import (
	"fmt"
	"log"
	"os"
    "encoding/json"
)

func main(){
    var mappings []interface{}

    entries, err := os.ReadDir("./commands")
    if err != nil {
        log.Fatal("no requests folder found")
    }


    for _, command := range entries {
        fmt.Println(command.Name())
    }

}
