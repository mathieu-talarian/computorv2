package computor

import (
	"fmt"
	"strings"
	"unicode"
)

var opChar = []byte{'*', '/', '%', '+', '-'}

type IParser interface {
	Start() (err error)
	CheckErrors() (err error)
	ParseOP() (err error)
	removeSpaces()
}

type Dt struct {
	Action string
	Op     interface{}
}
type Parser struct {
	Dt
	*CompErrors
	*Operandis
	input  string
	aInput []byte
}

func NewParser(input string) *Parser {
	i := strings.TrimSpace(input)
	return &Parser{
		CompErrors: NewComperrors(i),
		Operandis:  nil,
		input:      i,
		aInput:     []byte(i),
	}
}

func (p *Parser) Start() (err error) {
	p.removeSpaces()
	if err = p.TestQuestion(); err != nil {
		return err
	}
	if err = p.IllegalChar(); err != nil {
		return err
	}
	if i := p.TestEquals(); i == 0 {
		return p.parseforOne()
	} else if i == 1 {
		// do left right
	} else if i > 1 {
		return fmt.Errorf("Too much `=`")
	}
	return
}

func checkQuestionOne(a []byte) (err error) {
	for _, v := range a {
		if v == '?' {
			return fmt.Errorf("IllegalChar `?`")
		}
	}
	return
}

func isOperator(el byte) int {
	for _, v := range opChar {
		if el == v {
			return int(v)
		}
	}
	return 0
}

func constructOp(nt *[][]byte, i []byte) (err error) {
	for k, v := range i {
		if isOperator(v) != 0 {
			*nt = append(*nt, append(make([]byte, 0), i[:k]...))
			*nt = append(*nt, append(make([]byte, 0), i[k]))
			i = append(make([]byte, 0), i[k+1:]...)
			return constructOp(nt, i)
		}

	}
	*nt = append(*nt, append(make([]byte, 0), i...))
	return
}

func (p *Parser) parseforOne() (err error) {
	var nt = [][]byte{}
	if err = checkQuestionOne(p.aInput); err != nil {
		return
	}
	if err = constructOp(&nt, p.aInput); err != nil {
		return
	}
	p.Dt = Dt{
		Action: "OPERATION",
		Op:     nt,
	}
	return
}

/** ********************************************************************************************************* */
/** ********************************************************************************************************* */
/** ********************************************************************************************************* */
/** ********************************************************************************************************* */
func (p *Parser) removeSpaces() {
	for k, v := range p.aInput {
		if unicode.IsSpace(rune(v)) {
			p.aInput = append(p.aInput[:k], p.aInput[k+1:]...)
			p.removeSpaces()
		}
	}
}

func (p *Parser) TestEquals() int {
	var cpt = 0

	for _, el := range p.aInput {
		if el == '=' {
			cpt++
		}
	}
	return cpt
}
