package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) 
}

func Sqrt(x float64) (float64, ErrNegativeSqrt) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), 0
}

func printResult(x float64, e ErrNegativeSqrt) {
	if (e != 0)  {
		fmt.Println(e)
	} else {
		fmt.Println(x)
	}
}
func main() {
	printResult(Sqrt(2))
	printResult(Sqrt(-2))
}
