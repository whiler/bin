package bin

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

type BigEndianUnmarshaler interface {
	UnmarshalBigEndian(data []byte) (used int, err error)
}

type LittleEndianUnmarshaler interface {
	UnmarshalLittleEndian(data []byte) (used int, err error)
}

func UnmarshalBigEndian(input []byte, ins interface{}) error {
	return unmarshal(input, ins, binary.BigEndian, bigEndianUnmarshalerType, bigEndianUnmarshaler)
}

func UnmarshalLittleEndian(input []byte, ins interface{}) error {
	return unmarshal(input, ins, binary.LittleEndian, littleEndianUnmarshalerType, littleEndianUnmarshaler)
}

var (
	bigEndianUnmarshalerType    = reflect.TypeOf(new(BigEndianUnmarshaler)).Elem()
	littleEndianUnmarshalerType = reflect.TypeOf(new(LittleEndianUnmarshaler)).Elem()
)

type unmarshalerFunc func(interface{}, []byte) (int, error)

func bigEndianUnmarshaler(v interface{}, data []byte) (used int, err error) {
	unmarshaler := v.(BigEndianUnmarshaler)
	return unmarshaler.UnmarshalBigEndian(data)
}

func littleEndianUnmarshaler(v interface{}, data []byte) (used int, err error) {
	unmarshaler := v.(LittleEndianUnmarshaler)
	return unmarshaler.UnmarshalLittleEndian(data)
}

func unmarshal(input []byte, ins interface{}, order binary.ByteOrder, unmarshalerType reflect.Type, unmarshaler unmarshalerFunc) error {
	var (
		stack  valueStack
		cur    reflect.Value
		tpe    reflect.Type
		err    error
		offset int = 0
		delta  int = 0
		size   int = len(input)
	)

	cur = reflect.ValueOf(ins)
	if cur.Kind() != reflect.Ptr || cur.IsNil() {
		return fmt.Errorf("Invalid Unmarshal Type %s", reflect.TypeOf(ins))
	}

	stack.Push(cur)
	for len(stack) > 0 && err == nil {
		cur = stack.Pop()

		if !cur.IsValid() {
			err = fmt.Errorf("Invalid Value")
			break
		}

		tpe = cur.Type()
		if tpe.Implements(unmarshalerType) {
			delta, err = unmarshaler(cur.Interface(), input[offset:])
			offset += delta
			continue
		}

		switch kind := cur.Kind(); kind {
		case reflect.Ptr:
			if cur.IsNil() {
				cur.Set(reflect.New(tpe.Elem()))
			}
			stack.Push(cur.Elem())

		case reflect.Struct:
			err = handleStructKind(&cur, tpe, &stack)

		case reflect.Slice, reflect.Array:
			for i := cur.Len() - 1; i >= 0; i-- {
				stack.Push(cur.Index(i))
			}

		case reflect.Bool,
			reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128:
			delta = int(tpe.Size())
			if offset+delta <= size {
				setValue(&cur, kind, order, input[offset:offset+delta])
				offset += delta
			} else {
				err = fmt.Errorf("Need more %d byte(s)", offset+delta-size)
			}

		default:
			err = fmt.Errorf("Unsupported kind %s", kind)
		}
	}

	return err
}

func setValue(cur *reflect.Value, kind reflect.Kind, order binary.ByteOrder, data []byte) {
	switch kind {
	case reflect.Bool:
		cur.SetBool(data[0] != 0)

	case reflect.Int8:
		cur.SetInt(int64(data[0]))
	case reflect.Int16:
		cur.SetInt(int64(order.Uint16(data)))
	case reflect.Int32:
		cur.SetInt(int64(order.Uint32(data)))
	case reflect.Int64:
		cur.SetInt(int64(order.Uint64(data)))

	case reflect.Uint8:
		cur.SetUint(uint64(data[0]))
	case reflect.Uint16:
		cur.SetUint(uint64(order.Uint16(data)))
	case reflect.Uint32:
		cur.SetUint(uint64(order.Uint32(data)))
	case reflect.Uint64:
		cur.SetUint(order.Uint64(data))

	case reflect.Float32:
		cur.SetFloat(float64(math.Float32frombits(order.Uint32(data))))
	case reflect.Float64:
		cur.SetFloat(math.Float64frombits(order.Uint64(data)))

	case reflect.Complex64:
		cur.SetComplex(complex(
			float64(math.Float32frombits(order.Uint32(data[:4]))),
			float64(math.Float32frombits(order.Uint32(data[4:]))),
		))
	case reflect.Complex128:
		cur.SetComplex(complex(
			math.Float64frombits(order.Uint64(data[:8])),
			math.Float64frombits(order.Uint64(data[8:])),
		))
	}
}
