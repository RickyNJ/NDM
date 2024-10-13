package mocks

import (
	"encoding/json"
	"os"
)

type mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
}

func Generate(filepath string) *mock {
    var m mock 
    f, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }

    json.Unmarshal(f, &m)
    return &m 
}
