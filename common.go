package bin

import (
	"strconv"
)

const (
	TagName string = "bin"
)

func getIndexFromTag(tag string, i int) (int, error) {
	switch tag {
	case "":
		return i, nil

	case "-":
		return -1, nil

	default:
		if i64, e := strconv.ParseUint(tag, 10, strconv.IntSize); e != nil {
			return -1, e
		} else {
			return int(i64), e
		}
	}
}
