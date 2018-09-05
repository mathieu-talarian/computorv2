package main

import (
	"bufio"
	"os"
	"fmt"
)

type Prompt struct {
}

func NewPrompt(c chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf(">")
		c <- scanner.Text()
	}
}
