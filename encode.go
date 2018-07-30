package bin

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
)

var (
	bigEndianMarshalerType    = reflect.TypeOf(new(BigEndianMarshaler)).Elem()
	littleEndianMarshalerType = reflect.TypeOf(new(LittleEndianMarshaler)).Elem()
)

type BigEndianMarshaler interface {
	MarshalBigEndian() ([]byte, error)
}

type LittleEndianMarshaler interface {
	MarshalLittleEndian() ([]byte, error)
}

// {{{ valueStack
type valueStack []reflect.Value

func (stack *valueStack) Push(value reflect.Value) {
	*stack = append(*stack, value)
}

func (stack *valueStack) Pop() (ret reflect.Value) {
	size := len(*stack)
	ret, *stack = (*stack)[size-1], (*stack)[:size-1]
	return
}

// }}}

func MarshalBigEndian(ins interface{}) ([]byte, error) {
	var (
		buf   *bytes.Buffer = &bytes.Buffer{}
		stack valueStack    = valueStack{}
		cur   reflect.Value
		tpe   reflect.Type
		err   error
	)

	stack.Push(reflect.ValueOf(ins))
	for len(stack) > 0 && err == nil {
		cur = stack.Pop()

		if !cur.IsValid() {
			err = errors.New("Unexcepted error")
			break
		}

		tpe = cur.Type()
		if tpe.Implements(bigEndianMarshalerType) {
			marshaler := cur.Interface().(BigEndianMarshaler)
			if data, e := marshaler.MarshalBigEndian(); e != nil {
				err = e
				break
			} else if _, err = buf.Write(data); err != nil {
				break
			} else {
				continue
			}
		}

		switch kind := cur.Kind(); kind {
		case reflect.Ptr:
			if cur.IsNil() {
				stack.Push(reflect.New(tpe.Elem()))
			} else {
				stack.Push(cur.Elem())
			}

		case reflect.Struct:
			numField := cur.NumField()
			values := make([]reflect.Value, numField)
			idx := 0
			allowInvalid := true
			for i := numField - 1; i >= 0; i-- {
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
				case values[numField-idx-1].IsValid():
					err = errors.New("Field index duplicated")
					break
				default:
					values[numField-idx-1] = cur.Field(i)
				}
			}
			if err != nil {
				break
			}
			for _, value := range values {
				switch {
				case allowInvalid && value.IsValid():
					allowInvalid = false
					stack.Push(value)
				case !allowInvalid && !value.IsValid():
					err = errors.New("Field index is invalid")
					break
				case !allowInvalid && value.IsValid():
					stack.Push(value)
				}
				if err != nil {
					break
				}
			}

		case reflect.Slice, reflect.Array:
			for i := cur.Len() - 1; i >= 0; i-- {
				stack.Push(cur.Index(i))
			}

		case reflect.String:
			data := []byte(cur.Interface().(string))
			_, err = buf.Write(data)

		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128:
			err = binary.Write(buf, binary.BigEndian, cur.Interface())

		default:
			err = errors.New("Unsupported kind:" + string(kind))
		}
	}

	return buf.Bytes(), err
}
