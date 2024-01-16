package helpers

import (
	"errors"
	"strconv"
)

// Check if number is a valid one
func IsValidNumber(num string) (int, error) {
	if num == "" {
		return 0, errors.New("interger can not be empty string")
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	if n <= 0 {
		return 0, errors.New("integer can not be less or equal to zero")
	}
	return n, nil
}
