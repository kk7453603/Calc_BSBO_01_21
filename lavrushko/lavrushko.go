package lavrushko

import "math"

func Pow(number, power int) int {
	return int(math.Pow(float64(number), float64(power)))
}

func Sqrt(number, root int) int {
	return int(math.Sqrt(float64(number)))
}
