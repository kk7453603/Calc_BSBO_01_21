package negodin

func Mod(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func Left_Shift(a, b int) int {
	return a << b
}

func Right_Shift(a, b int) int {
	return a >> b
}
