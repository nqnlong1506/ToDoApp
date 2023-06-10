package api

import "strconv"

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
