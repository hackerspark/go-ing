package diffsquares

import "math"

func SquareOfSum(n int) int {
    squareOfSum := 0
	for i := 1; i <= n; i++ {
        squareOfSum += i
    }
    return int(math.Pow(float64(squareOfSum), 2))
}

func SumOfSquares(n int) int {
	sumOfSquares := 0
    for i := 1; i <= n; i++ {
        sumOfSquares += int(math.Pow(float64(i), 2))
    }
    return sumOfSquares
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
