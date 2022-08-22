package main

import (
	"hexagonalcalc/adapters/intcalc"
	"hexagonalcalc/adapters/storage/repositories"
	"hexagonalcalc/adapters/stringcalc"
	"hexagonalcalc/application"
)

func main() {
	resultRepo := repositories.NewResultRepository()
	calculatorService := application.NewCalculatorService(resultRepo)
	stringcalc.Execute(calculatorService, stringcalc.ExecuteParams{
		Left:      stringcalc.StringFour,
		Right:     stringcalc.StringThree,
		Operation: stringcalc.OperationAdd,
	})
	intcalc.Execute(calculatorService, intcalc.ExecuteParams{
		Left:      4,
		Right:     3,
		Operation: intcalc.OperationAdd,
	})
}
