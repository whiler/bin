package bin

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
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
		err   error
	)

	stack.Push(reflect.ValueOf(ins))
	for len(stack) > 0 {
		cur = stack.Pop()
		log.Println("current", cur)

		if !cur.IsValid() {
			continue
		}

		if cur.Type().Implements(bigEndianMarshalerType) {
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
				stack.Push(reflect.New(cur.Type().Elem()))
			} else {
				stack.Push(cur.Elem())
			}

		case reflect.Struct:
			for i := cur.NumField() - 1; i >= 0; i-- {
				stack.Push(cur.Field(i))
			}

		case reflect.Slice, reflect.Array:
			for i := cur.Len() - 1; i >= 0; i-- {
				stack.Push(cur.Index(i))
			}

		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.String:
			if err = binary.Write(buf, binary.BigEndian, cur.Interface()); err != nil {
				break
			}

		default:
			err = errors.New("Unsupported kind:" + string(kind))
			break
		}
	}

	return buf.Bytes(), err
}
