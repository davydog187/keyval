package main

import (
	"bufio"
	"fmt"
	"os"
)

func Add(a int, b int) int {
	return a + b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("got: " + scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
