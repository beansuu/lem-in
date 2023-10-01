package ant

import (
	"strconv"
)

func ParseNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
