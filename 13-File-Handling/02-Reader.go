package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("empty.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Println(reader)
}
