package conversion

import (
	"fmt"
	"strconv"
)

func StringToUint(strVal string) (uint, error) {
	uintVal64, err := strconv.ParseUint(strVal, 10, 64)
	uintVal := uint(uintVal64)
	return uintVal, err
}

func StringToFloat32(strVal string) (float32, error) {
	floatVal64, err := strconv.ParseFloat(strVal, 64)
	floatVal32 := float32(floatVal64)
	return floatVal32, err
}

func StringToFloat64(strVal string) (float64, error) {
	floatVal64, err := strconv.ParseFloat(strVal, 64)
	return floatVal64, err
}

func UintToString(uintVal uint) string {
	return strconv.Itoa(int(uintVal))
}

func Float32ToString(float32 float32) string {
	return fmt.Sprintf("%f", float32)
}
