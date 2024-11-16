package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
}

func main() {
    person := Person{"Alice", 30}
    data, _ := json.Marshal(person)
    fmt.Println(string(data)) // Output: {"name":"Alice","age":30}
}