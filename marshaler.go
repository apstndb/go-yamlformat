package yamlformat

import (
	"math"
	"strconv"
)

// marshalFloat64 formats float64 without scientific notation
func marshalFloat64(v float64) ([]byte, error) {
	// Special cases
	if math.IsNaN(v) {
		return []byte(".nan"), nil
	}
	if math.IsInf(v, 1) {
		return []byte(".inf"), nil
	}
	if math.IsInf(v, -1) {
		return []byte("-.inf"), nil
	}
	
	// Use strconv.FormatFloat with 'f' format to avoid scientific notation
	// -1 precision means use the smallest number of digits necessary
	str := strconv.FormatFloat(v, 'f', -1, 64)
	
	return []byte(str), nil
}