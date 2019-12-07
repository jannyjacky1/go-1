package hw_2

import (
	"errors"
	"strconv"
)

func UnpackString(str string) (string, error) {
	result := ""
	str += " "
	n := len(str)
	num := ""

	if _, err := strconv.Atoi(string(str[0])); err == nil {
		return "", errors.New("incorrect string")
	}

	for i := 1; i < n; i++ {
		symbol := string(str[i])
		if _, err := strconv.Atoi(symbol); err == nil {
			num += symbol
		} else {
			if cnt, err := strconv.Atoi(num); err == nil {
				symbol = string(str[i-len(num)-1])
				for i := 0; i < cnt; i++ {
					result += symbol
				}
				num = ""
			} else {
				result += string(str[i-1])
			}
		}
	}
	return result, nil
}
