package model

import (
	"fmt"
	"strconv"
)

func Round8(x float64) float64 {
	s := fmt.Sprintf("%.8f", x)
	y, _ := strconv.ParseFloat(s, 64)
	return y
}

func Round4(x float64) float64 {
	s := fmt.Sprintf("%.4f", x)
	y, _ := strconv.ParseFloat(s, 64)
	return y
}

func Float64Round(x float64, prec ...int) float64 {
	precision := 4
	if len(prec) == 1 {
		precision = prec[0]
	}
	format := "%." + strconv.Itoa(precision) + "f"
	s := fmt.Sprintf(format, x)
	y, _ := strconv.ParseFloat(s, 64)
	return y
}
func StringToFloat64(s string) float64 {
	y, _ := strconv.ParseFloat(s, 64)
	return y
}
func Float64ToString(f float64) string {
	return fmt.Sprint(f)
}
