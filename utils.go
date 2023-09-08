package mgom

import "strconv"

func StringToBool(value string) bool {
	p, _ := strconv.ParseBool(value)
	return p
}

func StringToInt(value string) int {
	p, _ := strconv.ParseInt(value, 10, 64)
	return int(p)
}
