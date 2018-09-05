package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type IComputor interface {
	Start(operandis Operandis)
}

const (
	LEFT = iota
	RIGHT
)

type Computor struct {
	Action    string
	Errors    []string
	LComputed interface{}
	RComputed interface{}
}

func (c *Computor) Start(operandis Operandis) {
	c.Errors = []string{}
	c.CheckAndCompute(LEFT, operandis.left)
	c.CheckAndCompute(RIGHT, operandis.right)
}

func (c *Computor) CheckAndCompute(side int, op []byte) {
	switch side {
	case LEFT:
		c.checkandcomputeLeft(op)
	case RIGHT:
		c.checkandcomputeRight(op)
	}
}

func (c *Computor) checkandcomputeLeft(op []byte) interface{} {
	if isNumber(string(op)) {
		fmt.Println("is Number")
		fmt.Println(strconv.ParseFloat(string(op), 10))
	} else {
		fmt.Println("is not number")
	}
	return nil
}

func (c *Computor) checkandcomputeRight(op []byte) interface{} {
	if len(op) == 1 && op[0] == '?' {
		c.Action = QUESTION
	}
	fmt.Println(string(op))
	return nil
}

func isNumber(test string) bool {
	var re = regexp.MustCompile(`(?m)^-?(\d+\.?\d*)$|(\d*\.?\d+)$`).FindStringSubmatch(test)
	if len(re) > 1 && re[1] != "" {
		return true
	}
	return false
}
