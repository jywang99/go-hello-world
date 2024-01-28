package interfaces

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	Number float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.Number)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt{x}
	}
	return x, nil
}

func TestError() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

