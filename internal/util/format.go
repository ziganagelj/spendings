package util

import (
	"fmt"
	"strconv"
)

var DateFormat = "2006-01-02"

// FormatFloat formats a float64 to a string with two decimal places.
//
// This function is useful for converting floating-point numbers to a
func FormatFloat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// ParseStringToUint parses a string into a uint.
// It returns an error if the string cannot be parsed or the value is out of range.
func ParseStringToUint(s string) (uint, error) {
	val, err := strconv.ParseUint(s, 10, 64) // Using 0 bit size to fit the uint size of the underlying platform
	return uint(val), err
}
