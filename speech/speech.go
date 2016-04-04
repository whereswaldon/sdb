package speech

import (
	"fmt"
	"github.com/VerbalExpressions/GoVerbalExpressions"
	"os"
	"os/exec"
	"strings"
)

/**
Speak exits if given empty string, and otherways says the utterance
aloud using the flite cli.
*/
func Speak(utterance string) {
	if utterance == "" {
		farewell := exec.Command("flite", "-t", "Goodbye")
		farewell.Run()
		os.Exit(0)
	}
	fmt.Printf("response: %s\n", utterance)
	flite := exec.Command("flite", "-t", utterance)
	err := flite.Run()
	if err != nil {
		fmt.Println(err)
	}
}

/**
Compose accepts an input string and returns a response
to that user input in the form of a string (or nil,
if the user asked to exit)
*/
func Compose(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "exit" || input == "quit" {
		return ""
	}

	v := verbalexpressions.
		New().
		StartOfLine().
		Anything().
		Then("have a").
		BeginCapture().
		Word().
		EndCapture().
		Anything().
		Captures(input)

	fmt.Printf("%s\n, %s", v, input)
	return input
}
