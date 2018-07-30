package bin

import (
	"errors"
	"reflect"
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

type valueStack []reflect.Value

func (stack *valueStack) Push(value reflect.Value) {
	*stack = append(*stack, value)
}

func (stack *valueStack) Pop() (ret reflect.Value) {
	size := len(*stack)
	ret, *stack = (*stack)[size-1], (*stack)[:size-1]
	return
}

func handleStructKind(cur *reflect.Value, tpe reflect.Type, stack *valueStack) (err error) {
	numField := cur.NumField()
	fields := make([]reflect.Value, numField)
	idx := 0
	allowInvalid := true

	for i := 0; i < numField && err == nil; i++ {
		tag := tpe.Field(i).Tag.Get(TagName)
		idx, err = getIndexFromTag(tag, i)
		switch {
		case err != nil:
			break
		case idx == -1:
			continue
		case idx >= numField:
			err = errors.New("Field index out of range")
			break
		case fields[numField-idx-1].IsValid():
			err = errors.New("Field index duplicated")
			break
		default:
			fields[numField-idx-1] = cur.Field(i)
		}
	}

	if err == nil {
		for _, value := range fields {
			switch {
			case allowInvalid && value.IsValid():
				allowInvalid = false
				stack.Push(value)
			case !allowInvalid && !value.IsValid():
				err = errors.New("Field index is invalid")
			case !allowInvalid && value.IsValid():
				stack.Push(value)
			}
			if err != nil {
				break
			}
		}
	}

	return
}
