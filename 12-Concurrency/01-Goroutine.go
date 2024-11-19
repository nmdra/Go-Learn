package main

import (
	"fmt"
	"time"
)

func speak(arg string) {
	fmt.Println(arg)
}

func main() {
	go speak("Hello World")
	fmt.Println("Hii")
	time.Sleep(500 * time.Millisecond)
}
