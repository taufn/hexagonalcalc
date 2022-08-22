package domain

type Calculator interface {
	Add(left, right int) (Result, error)
	Subtract(left, right int) (Result, error)
	Multiply(left, right int) (Result, error)
	Divide(left, right int) (Result, error)
}

type ResultRepository interface {
	Save(result Result) (Result, error)
}
