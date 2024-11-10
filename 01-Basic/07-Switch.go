package main

import "fmt"

func switchGo2() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}

	// Shorthand Declaration

	switch day := "monday"; day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}

	// We can also use the fallthrough keyword to transfer control to the next case even though the current case might have matched

	switch day := "monday"; day {
	case "monday":
		fmt.Println("time to work!")
		fallthrough
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}

}
