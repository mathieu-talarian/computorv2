package computor

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
)

type Computor struct {
	Variables *list.List
}

func NewComputor() *Computor {
	return &Computor{}
}

func (c *Computor) Exe(dt Dt) (err error) {
	if dt.Action == "OPERATION" {
		return doOp(dt.Op.([][]byte))
	}
	return nil
}

/** check with regexp if the string in param is a number, works with float */
func isNumber(test string) bool {
	var re = regexp.MustCompile(`(?m)^-?(\d+\.?\d*)$|(\d*\.?\d+)$`).FindStringSubmatch(test)
	if len(re) > 1 && re[1] != "" {
		return true
	}
	return false
}

func wrongVarChars(el byte) (err error) {
	var wrongvarchars = []byte{'i', 'x', 'y', 'e', 'f'}
	for _, v := range wrongvarchars {
		if v == el {
			return fmt.Errorf("Reserved variable character: `%c`", el)
		}
	}
	return
}

func isValidVar(v string, len int) (err error) {
	fmt.Println(len)
	if len == 1 {
		return wrongVarChars(v[0])
	}
	return
}

/** find is []byte in arg => if len == 1 && is operator and if is key is not the last or the first */
func checkCase(k int, v []byte, len int) (err error) {
	if len != 0 {
		if len == 1 {
			if op := isOperator(v[0]); op != 0 {
				return // operator and no errors
			} else if isNumber(string(v)) {
				return // number and no errors
			} else if err = isValidVar(string(v), 1); err != nil {
				return // var and no errors
			} else {
				return err
			}
		}
		if isNumber(string(v)) {
			_, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}
		} else if err = isValidVar(string(v), len); err != nil {
			return // var and no errors
		} else {
			return err
		}
	}
	return
}

func doOp(dt [][]byte) (err error) {
	for k, v := range dt {
		fmt.Println(string(v))
		if err = checkCase(k, v, len(v)); err != nil {
			return
		}
	}
	return

}
