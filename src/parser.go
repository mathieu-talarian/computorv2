package computor

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
	ParseOP() (err error)
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
	if err = p.buildObj(p.findEqual()); err != nil {
		return
	}
	if err = checkQuestionLeft(p.Left); err != nil {
		return
	}

	return
}

func checkQuestionLeft(op []byte) (err error) {
	for _, v := range op {
		if v == '?' {
			return fmt.Errorf("Question mark on wrong side")
		}
	}
	return
}

func (p *Parser) buildObj(k int) (err error) {
	if k != 0 && k < len(p.aInput)-1 {
		p.Operandis.Left = p.aInput[0:k]
		p.Operandis.Right = p.aInput[k+1:]
		return
	}
	return fmt.Errorf("Equal in wrong spot")
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
		return fmt.Errorf("Not enough arguments")
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

func (p *Parser) ParseOp() (err error) {
	if err = p.Operandis.Parse(); err != nil {
		return
	}
}
