package computor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Main() {
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
			fmt.Println("\x1b[31m", err, "\x1b[0m")
		} else if err := parser.Start(); err != nil {
			fmt.Println("\x1b[31m", err, "\x1b[0m")
		} else if err := parser.ParseOp(); err != nil {
			fmt.Println("\x1b[31m", err, "\x1b[0m")
		}
	}
}
