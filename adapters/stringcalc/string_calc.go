package stringcalc

import (
	"hexagonalcalc/domain"
	"log"
)

type ExecuteParams struct {
	Left      Number
	Right     Number
	Operation string
}

const (
	OperationAdd      = "add"
	OperationSubtract = "subtract"
	OperationMultiply = "multiply"
	OperationDivide   = "divide"
)

func Execute(calculator domain.Calculator, params ExecuteParams) {
	var result domain.Result
	var err error

	switch params.Operation {
	case OperationAdd:
		result, err = calculator.Add(params.Left.toInt(), params.Right.toInt())
	case OperationSubtract:
		result, err = calculator.Subtract(params.Left.toInt(), params.Right.toInt())
	case OperationMultiply:
		result, err = calculator.Multiply(params.Left.toInt(), params.Right.toInt())
	case OperationDivide:
		result, err = calculator.Divide(params.Left.toInt(), params.Right.toInt())
	default:
		log.Fatalf("operation not supported: %s", params.Operation)
	}

	if err != nil {
		log.Panic(err)
	}
	log.Printf(
		"%s %s %s equals %s",
		params.Left.ToString(),
		params.Operation,
		params.Right.ToString(),
		ToNumber(result.Result).ToString(),
	)
}
