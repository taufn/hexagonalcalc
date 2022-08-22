package intcalc

import (
	"hexagonalcalc/domain"
	"log"
)

type ExecuteParams struct {
	Left      int
	Right     int
	Operation string
}

const (
	OperationAdd      = "tambah"
	OperationSubtract = "kurang"
	OperationMultiply = "kali"
	OperationDivide   = "bagi"
)

func Execute(calculator domain.Calculator, params ExecuteParams) {
	var result domain.Result
	var err error

	switch params.Operation {
	case OperationAdd:
		result, err = calculator.Add(params.Left, params.Right)
	case OperationSubtract:
		result, err = calculator.Subtract(params.Left, params.Right)
	case OperationMultiply:
		result, err = calculator.Multiply(params.Left, params.Right)
	case OperationDivide:
		result, err = calculator.Divide(params.Left, params.Right)
	default:
		log.Fatalf("operation not supported: %s", params.Operation)
	}

	if err != nil {
		log.Panic(err)
	}
	log.Printf(
		"%d %s %d sama dengan %d",
		params.Left,
		params.Operation,
		params.Right,
		result.Result,
	)
}
