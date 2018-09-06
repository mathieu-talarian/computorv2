package computor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Main() {
	reader := bufio.NewReader(os.Stdin)
	//computor := new(Computor)
	fmt.Println("Computor v2")
	fmt.Println("---------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		parser := NewParser(text)
		if err := parser.Start(); err != nil {
			fmt.Println(err)
		}
	}
}
