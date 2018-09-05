package main

import "fmt"

var illegalchars = []byte{
	'\'',
	'"',
}

type ICompErrors interface {
	Constructor(string)
	TestEquals() bool
	TestQuestion() bool
	IllegalChar()
}

type CompErrors struct {
	data   []string
	input  string
	aInput []byte
}

func (c *CompErrors) Constructor(prompt string) {
	c.input = prompt
	c.aInput = []byte(prompt)
	c.data = []string{}
}

func (c *CompErrors) TestEquals() bool {
	var cpt = 0

	for _, el := range c.aInput {
		if el == '=' {
			cpt++
		}
	}
	if cpt == 0 {
		c.data = append(c.data, errorRet('=', 0))
		return false
	}
	if cpt > 1 {
		c.data = append(c.data, errorRet('=', 1))
		return false
	}
	return true
}

func (c *CompErrors) TestQuestion() bool {
	var cpt = 0

	for _, el := range c.aInput {
		if el == '?' {
			cpt++
		}
	}
	if cpt > 1 {
		c.data = append(c.data, errorRet('?', 1))
		return false
	}
	return true
}

func (c *CompErrors) IllegalChar() {
	for _, el := range c.aInput {
		for _, el2 := range illegalchars {
			if el == el2 {
				c.data = append(c.data, errorRet(byte(el), 2))
			}
		}
	}
}

func errorRet(char byte, q int) string {
	switch q {
	case 0:
		return fmt.Sprintf("Not enough `%c`", char)
	case 1:
		return fmt.Sprintf("Too much `%c", char)
	case 2:
		return fmt.Sprintf("Illegal character: %c", char)
	}
	return ""
}
