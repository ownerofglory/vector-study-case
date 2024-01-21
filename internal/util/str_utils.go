package util

import (
	"fmt"
	"strconv"
	"strings"
)

func StrToIntSlice(s string) ([]int, error) {
	sList := strings.Split(s, ",")
	ints := make([]int, 0, len(sList))
	for _, str := range sList {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("unable to convert input string '%s' to []int: %v", s, err)
		}
		ints = append(ints, i)
	}

	return ints, nil
}
