package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Constants for signaling the end of the game and end of initialization
const (
	END_SIGNAL      = "END"
	END_INIT_SIGNAL = "END_INIT"
)

// PostMessage converts the given message to JSON format and prints it to standard output.
func PostMessage(message map[string]interface{}) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON message: %v\n", err)
		return
	}

	fmt.Println(string(jsonMessage))
}

// ReadMessage reads a line of input from standard input, parses it as JSON, and returns the parsed message.
func ReadMessage() interface{} {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return nil
	}

	var message interface{}
	err = json.Unmarshal([]byte(input), &message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling JSON message: %v\n", err)
		return nil
	}

	return message
}
