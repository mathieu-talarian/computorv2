package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type IParser interface {
	Constructor(string)
	Start() (err error)
	CheckErrors() (err error)
	removeSpaces()
}

type Parser struct {
	*CompErrors
	*Operandis
	input  string
	aInput []byte
}

func (p *Parser) Constructor(prompt string) {
	p.input = strings.TrimSpace(prompt)
	p.aInput = []byte(prompt)
	p.CompErrors = new(CompErrors)
	p.Operandis = new(Operandis)
}

func (p *Parser) Start() (err error) {
	p.removeSpaces()
	return p.buildObj(p.findEqual())
}

func (p *Parser) buildObj(k int) (err error) {
	if k != 0 && k < len(p.aInput)-1 {
		p.Operandis.left = p.aInput[0:k]
		p.Operandis.right = p.aInput[k+1:]
		return
	}
	return errors.New(fmt.Sprintf("Equal in wrong spot"))
}

func (p *Parser) findEqual() int {
	for k, v := range p.aInput {
		if v == '=' {
			return k
		}
	}
	return 0
}

func (p *Parser) CheckErrors() (err error) {
	if len(p.input) < 3 {
		return errors.New(fmt.Sprintf("Not enough arguments"))
	}
	p.CompErrors.Constructor(p.input)
	p.TestEquals()
	p.TestQuestion()
	p.IllegalChar()
	if len(p.data) > 0 {
		return errors.New(strings.Join(p.data, ", "))
	}
	return nil
}

func (p *Parser) removeSpaces() {
	for k, v := range p.aInput {
		if unicode.IsSpace(rune(v)) {
			p.aInput = append(p.aInput[:k], p.aInput[k+1:]...)
			p.removeSpaces()
		}
	}
}
