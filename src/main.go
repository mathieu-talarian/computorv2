package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	computor := new(Computor)
	fmt.Println("Computor v2")
	fmt.Println("---------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		parser := new(Parser)
		parser.Constructor(text)
		if err := parser.CheckErrors(); err != nil {
			fmt.Println(err)
		} else if err := parser.Start(); err != nil {
			fmt.Println(err)
		} else {
			computor.Start(*parser.Operandis)
		}
	}

}
