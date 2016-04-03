package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	incoming := make(chan string)
	outgoing := make(chan string)
	go listen(incoming)
	go compose(incoming, outgoing)
	speak(outgoing)
}

/**
listen continuously checks for new input and sends it
deeper in the application for processing
*/
func listen(toProcessing chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		toProcessing <- text
	}
}

/**
compose continuously processes input on the toProcessing
channel and sends responses out on the fromProcessing
channel
*/
func compose(toProcessing, fromProcessing chan string) {
	for {
		//for now, immediately pass on input
		input := <-toProcessing
		//fmt.Printf("input: %s", input)
		fromProcessing <- input
	}
}

/**
speak continuously listens for output and sends it to the
user.
*/
func speak(fromProcessing chan string) {
	for {
		prompt()
		fmt.Println(<-fromProcessing)
	}
}

/**
prompt shows the user a prompt to indicate that the program
is ready for input.
*/
func prompt() {
	fmt.Print(">")
}
