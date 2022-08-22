package domain

import (
	"fmt"
)

type Math struct {
	Left  int
	Right int
}

func NewMath(left, right int) (Math, error) {
	if left < numMin || left > numMax || right < numMin || right > numMax {
		return Math{}, fmt.Errorf("left %v or right %v number is out of bounds", left, right)
	}
	return Math{Left: left, Right: right}, nil
}

func (m Math) Add() int {
	return m.Left + m.Right
}

func (m Math) Subtract() int {
	return m.Left - m.Right
}

func (m Math) Multiply() int {
	return m.Left * m.Right
}

func (m Math) Divide() int {
	return m.Left / m.Right
}
