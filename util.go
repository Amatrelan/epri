package main

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint[K any](data K) {
	s, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(s))
}
