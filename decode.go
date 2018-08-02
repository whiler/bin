package bin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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
	return unmarshal(bytes.NewBuffer(input), ins, binary.BigEndian, bigEndianUnmarshalerType, bigEndianUnmarshaler)
}

func UnmarshalLittleEndian(input []byte, ins interface{}) error {
	return unmarshal(bytes.NewBuffer(input), ins, binary.LittleEndian, littleEndianUnmarshalerType, littleEndianUnmarshaler)
}

func UnmarshalBigEndianFrom(reader io.Reader, ins interface{}) error {
	return unmarshal(reader, ins, binary.BigEndian, bigEndianUnmarshalerType, bigEndianUnmarshaler)
}

func UnmarshalLittleEndianFrom(reader io.Reader, ins interface{}) error {
	return unmarshal(reader, ins, binary.LittleEndian, littleEndianUnmarshalerType, littleEndianUnmarshaler)
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

func unmarshal(reader io.Reader, ins interface{}, order binary.ByteOrder, unmarshalerType reflect.Type, unmarshaler unmarshalerFunc) error {
	var (
		stack      valueStack
		cur        reflect.Value
		tpe        reflect.Type
		kind       reflect.Kind
		err        error
		buf        []byte
		backReader = newBackfillReader(reader)
		size       int
		delta      int
	)

	cur = reflect.ValueOf(ins)
	if cur.Kind() != reflect.Ptr || cur.IsNil() {
		return fmt.Errorf("Invalid Unmarshal Type %#v", ins)
	}

	stack.Push(cur)
	for len(stack) > 0 && err == nil {
		cur = stack.Pop()

		tpe = cur.Type()
		kind = cur.Kind()
		if tpe.Implements(unmarshalerType) {
			if kind == reflect.Ptr && cur.IsNil() {
				cur.Set(reflect.New(tpe.Elem()))
			}
			buf = make([]byte, defaultBufSize)
			if size, err = backReader.Read(buf); err != nil {
				continue
			}
			if delta, err = unmarshaler(cur.Interface(), buf); err == nil {
				_, err = backReader.Backfill(buf[delta:size])
			}
			continue
		}

		switch kind {
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
			buf = make([]byte, tpe.Size())
			if _, err = io.ReadFull(backReader, buf); err == nil {
				setValue(&cur, kind, order, buf)
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
			float64(math.Float32frombits(order.Uint32(data[0:4]))),
			float64(math.Float32frombits(order.Uint32(data[4:8]))),
		))
	case reflect.Complex128:
		cur.SetComplex(complex(
			math.Float64frombits(order.Uint64(data[0:8])),
			math.Float64frombits(order.Uint64(data[8:16])),
		))
	}
}
