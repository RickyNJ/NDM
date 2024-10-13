package mocks

import (
	"encoding/json"
	"fmt"
	"os"
)

type mock struct {
    Command string `json:"command"`
    Response string `json:"response"`
    Arguments string `json:"arguments"`
    Flags string `json:"flags"`
}

func GenerateFromJSON(filepath string) *mock {
    m := &mock{} 
    f, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }

    unmarshallingError := json.Unmarshal(f, m)
    if unmarshallingError != nil {
        panic("error while mashaling ")
    }

    return m 
}

func ReadMappingsDir() []*mock {
    var result []*mock
    mappings, err  := os.ReadDir("mappings")
    if err != nil{
        panic("error while opening mappings directory")
    }

    for _, v := range mappings {
        var mapname string
        mapname = "mappings/" + v.Name()
        fmt.Println(mapname)
        m := GenerateFromJSON(mapname)
        result = append(result, m)
    }

    return result
}
