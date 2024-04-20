package main

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	var tests = []struct {
		line    string
		command Command
	}{
		{line: "READ a", command: Read{key: "a"}},
		{line: "WRITE a b", command: Write{key: "a", val: "b"}},
		{line: "DELETE a", command: Delete{key: "a"}},
	}

	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			if result, err := ParseCommand(test.line); result != test.command {
				if err != nil {
					t.Errorf("received error %v", err)
				}

				t.Errorf("expected %+v to be %+v got %+v", test.line, test.command, result)
			}
		})
	}

}

func TestExecuteReadCommand(t *testing.T) {
	var tests = []struct {
		state   State
		command Command
		message string
	}{
		{state: State{Data: map[string]string{}}, command: Read{key: "a"}, message: "Key not found: a"},
		{state: State{Data: map[string]string{"a": "1"}}, command: Read{key: "a"}, message: "1"},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			message, err := test.command.ExecuteCommand(test.state)

			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			if message != test.message {
				t.Errorf("expected %s but got %s", test.message, message)
			}
		})
	}
}
