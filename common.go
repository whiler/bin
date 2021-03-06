package bin

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

const (
	// TagName defines the especial tag name in struct field we are using
	TagName        string = "bin"
	defaultBufSize        = 4096
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
	ret, *stack = (*stack)[size-1], (*stack)[0:size-1]
	return
}

func handleStructKind(cur *reflect.Value, tpe reflect.Type, stack *valueStack) (err error) {
	var (
		numField     = cur.NumField()
		fields       = make([]reflect.Value, numField)
		idx          int
		allowInvalid = true
	)

	for i := 0; i < numField && err == nil; i++ {
		tag := tpe.Field(i).Tag.Get(TagName)
		idx, err = getIndexFromTag(tag, i)
		switch {
		case err != nil:
			break
		case idx == -1:
			continue
		case idx >= numField:
			err = fmt.Errorf("Field index '%d' out of range", idx)
			break
		case fields[numField-idx-1].IsValid():
			err = fmt.Errorf("Field index '%d' duplicated", idx)
			break
		default:
			fields[numField-idx-1] = cur.Field(i)
		}
	}

	if err == nil {
		for _, field := range fields {
			switch {
			case allowInvalid && field.IsValid():
				allowInvalid = false
				stack.Push(field)
			case !allowInvalid && !field.IsValid():
				err = fmt.Errorf("Field indexes invalid")
			case !allowInvalid && field.IsValid():
				stack.Push(field)
			}
			if err != nil {
				break
			}
		}
	}

	return
}

type backfillReader struct {
	buffer *bytes.Buffer
	reader io.Reader
}

func newBackfillReader(reader io.Reader) *backfillReader {
	return &backfillReader{new(bytes.Buffer), reader}
}

func (backReader *backfillReader) Read(dst []byte) (int, error) {
	if size, _ := backReader.buffer.Read(dst); size > 0 {
		return size, nil
	}
	return backReader.reader.Read(dst)
}

func (backReader *backfillReader) Backfill(src []byte) (n int, err error) {
	return backReader.buffer.Write(src)
}
