package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type State struct {
	Data map[string]string
}

type Command interface {
	ExecuteCommand(State) (string, error)
}

type Read struct {
	key string
}

type Write struct {
	key string
	val string
}

type Delete struct {
	key string
}

func (c Read) ExecuteCommand(state State) (string, error) {
	if val, ok := state.Data[c.key]; ok != false {
		return val, nil
	}

	return fmt.Sprintf("Key not found: %s", c.key), nil
}

func (c Write) ExecuteCommand(state State) (string, error) {
	state.Data[c.key] = c.val

	return "", nil
}

func (c Delete) ExecuteCommand(state State) (string, error) {
	delete(state.Data, c.key)

	return "", nil
}

func ParseCommand(line string) (Command, error) {
	tokens := strings.Split(line, " ")

	switch strings.ToLower(tokens[0]) {
	case "read":
		if len(tokens) != 2 {
			return nil, errors.New("READ must have a key")
		}

		key := tokens[1]
		return Read{key}, nil
	case "write":
		if len(tokens) != 3 {
			return nil, errors.New("WRITE must have a key and value")
		}
		key := tokens[1]
		val := tokens[2]
		return Write{key, val}, nil
	case "delete":
		if len(tokens) != 2 {
			return nil, errors.New("DELETE must have a key")
		}

		key := tokens[1]
		return Delete{key}, nil

	default:
		return nil, errors.New("Failed to parse command")
	}
}

func main() {
	state := State{Data: map[string]string{}}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		command, err := ParseCommand(line)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

		message, err := command.ExecuteCommand(state)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to execute command %v because %s\n", command, err)
			continue
		}

		if message != "" {
			fmt.Println(message)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
