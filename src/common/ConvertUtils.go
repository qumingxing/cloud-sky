package common

import (
	"strconv"
)

func StringToInt(val string) int {
	res, _ := strconv.Atoi(val)
	return res
}
func IntToString(val int) string {
	res := strconv.Itoa(val)
	return res
}
func StringToFloat(val string) float64 {
	res, _ := strconv.ParseFloat(val, 0)
	return res
}
