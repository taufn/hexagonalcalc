package stringcalc

type Number string

const (
	StringZero  Number = "zero"
	StringOne          = "one"
	StringTwo          = "two"
	StringThree        = "three"
	StringFour         = "four"
	StringFive         = "five"
	StringSix          = "six"
	StringSeven        = "seven"
	StringEight        = "eight"
	StringNine         = "nine"
)

const (
	NumZero  = 0
	NumOne   = 1
	NumTwo   = 2
	NumThree = 3
	NumFour  = 4
	NumFive  = 5
	NumSix   = 6
	NumSeven = 7
	NumEight = 8
	NumNine  = 9
)

func ToNumber(n int) Number {
	switch n {
	case NumNine:
		return StringNine
	case NumEight:
		return StringEight
	case NumSeven:
		return StringSeven
	case NumSix:
		return StringSix
	case NumFive:
		return StringFive
	case NumFour:
		return StringFour
	case NumThree:
		return StringThree
	case NumTwo:
		return StringTwo
	case NumOne:
		return StringOne
	case NumZero:
	default:
	}
	return StringZero
}

func (n Number) toInt() int {
	switch n {
	case StringNine:
		return NumNine
	case StringEight:
		return NumEight
	case StringSeven:
		return NumSeven
	case StringSix:
		return NumSix
	case StringFive:
		return NumFive
	case StringFour:
		return NumFour
	case StringThree:
		return NumThree
	case StringTwo:
		return NumTwo
	case StringOne:
		return NumOne
	case StringZero:
	default:
	}
	return NumZero
}

func (n Number) ToString() string {
	return string(n)
}
