package computor

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

func NewComperrors(input string) *CompErrors {
	return &CompErrors{
		input:  input,
		aInput: []byte(input),
		data:   []string{},
	}
}

func (c *CompErrors) TestQuestion() error {
	var cpt = 0

	for _, el := range c.aInput {
		if el == '?' {
			cpt++
		}
	}
	if cpt > 1 {
		return fmt.Errorf("Too much `?`")
	}
	return nil
}

func (c *CompErrors) IllegalChar() (err error) {
	for _, el := range c.aInput {
		for _, el2 := range illegalchars {
			if el == el2 {
				return fmt.Errorf("Illegal character: `%c`", el)
			}
		}
	}
	return
}

/** ********************************************************************************************************* */
/** ********************************************************************************************************* */
/** ********************************************************************************************************* */
/** ********************************************************************************************************* */
