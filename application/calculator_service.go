package application

import "hexagonalcalc/domain"

type CalculatorService struct {
	resultRepo domain.ResultRepository
}

func NewCalculatorService(resultRepo domain.ResultRepository) CalculatorService {
	return CalculatorService{resultRepo: resultRepo}
}

func (s CalculatorService) Add(left, right int) (result domain.Result, err error) {
	mathObj, err := domain.NewMath(left, right)
	if err != nil {
		return
	}
	return s.resultRepo.Save(domain.Result{
		Left:   left,
		Right:  right,
		Result: mathObj.Add(),
	})
}

func (s CalculatorService) Subtract(left, right int) (result domain.Result, err error) {
	mathObj, err := domain.NewMath(left, right)
	if err != nil {
		return
	}
	return s.resultRepo.Save(domain.Result{
		Left:   left,
		Right:  right,
		Result: mathObj.Subtract(),
	})
}

func (s CalculatorService) Multiply(left, right int) (result domain.Result, err error) {
	mathObj, err := domain.NewMath(left, right)
	if err != nil {
		return
	}
	return s.resultRepo.Save(domain.Result{
		Left:   left,
		Right:  right,
		Result: mathObj.Multiply(),
	})
}

func (s CalculatorService) Divide(left, right int) (result domain.Result, err error) {
	mathObj, err := domain.NewMath(left, right)
	if err != nil {
		return
	}
	return s.resultRepo.Save(domain.Result{
		Left:   left,
		Right:  right,
		Result: mathObj.Multiply(),
	})
}
