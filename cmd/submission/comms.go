package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	END_SIGNAL      = "END"
	END_INIT_SIGNAL = "END_INIT"
)

func PostMessage(message map[string]interface{}) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON message: %v\n", err)
		return
	}

	fmt.Println(string(jsonMessage))
}

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
