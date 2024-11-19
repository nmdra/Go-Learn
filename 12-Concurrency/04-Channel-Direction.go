package main

import "fmt"

// ping function sends a message to the pings channel.
// The `chan<- string` type specifies that this channel is send-only.
func ping(pings chan<- string, msg string) {
	pings <- msg // Send the message into the pings channel.
}

// pong function receives a message from the pings channel (receive-only)
// and sends it to the pongs channel (send-only).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // Receive the message from the pings channel.
	pongs <- msg   // Send the message into the pongs channel.
}

func main() {
	// Create a buffered channel `pings` with capacity 1, used for sending messages.
	pings := make(chan string, 1)

	// Create another buffered channel `pongs` with capacity 1, used for forwarding messages.
	pongs := make(chan string, 1)

	// Call the ping function to send a message to the `pings` channel.
	ping(pings, "passed message")

	// Call the pong function to forward the message from `pings` to `pongs`.
	pong(pings, pongs)

	// Receive and print the message from the `pongs` channel.
	fmt.Println(<-pongs)
}
