package pkg

import (
	"strconv"
	"strings"
)

func ConvertStrToInt64Slice(s string) ([]int64, error) {
	trimmed := strings.Trim(s, "[]")
	stringsSlice := strings.Split(trimmed, ", ")
	newInts := make([]int64, len(stringsSlice))

	for i, str := range stringsSlice {
		newInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		newInts[i] = newInt
	}
	return newInts, nil
}
