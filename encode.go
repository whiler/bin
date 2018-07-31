package bin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type BigEndianMarshaler interface {
	MarshalBigEndian() ([]byte, error)
}

type LittleEndianMarshaler interface {
	MarshalLittleEndian() ([]byte, error)
}

func MarshalBigEndian(ins interface{}) ([]byte, error) {
	return marshal(ins, binary.BigEndian, bigEndianMarshalerType, bigEndianMarshaler)
}

func MarshalLittleEndian(ins interface{}) ([]byte, error) {
	return marshal(ins, binary.LittleEndian, littleEndianMarshalerType, littleEndianMarshaler)
}

var (
	bigEndianMarshalerType    = reflect.TypeOf(new(BigEndianMarshaler)).Elem()
	littleEndianMarshalerType = reflect.TypeOf(new(LittleEndianMarshaler)).Elem()
)

type marshalerFunc func(v interface{}) ([]byte, error)

func bigEndianMarshaler(v interface{}) ([]byte, error) {
	marshaler := v.(BigEndianMarshaler)
	return marshaler.MarshalBigEndian()
}

func littleEndianMarshaler(v interface{}) ([]byte, error) {
	marshaler := v.(LittleEndianMarshaler)
	return marshaler.MarshalLittleEndian()
}

func marshal(ins interface{}, order binary.ByteOrder, marshalerType reflect.Type, marshaler marshalerFunc) ([]byte, error) {
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
			err = fmt.Errorf("Unexcepted error")
			break
		}

		tpe = cur.Type()
		if tpe.Implements(marshalerType) {
			if data, e := marshaler(cur.Interface()); e != nil {
				err = e
			} else {
				_, err = buf.Write(data)
			}
			continue
		}

		switch kind := cur.Kind(); kind {
		case reflect.Ptr:
			if cur.IsNil() {
				stack.Push(reflect.New(tpe.Elem()))
			} else {
				stack.Push(cur.Elem())
			}

		case reflect.Struct:
			err = handleStructKind(&cur, tpe, &stack)

		case reflect.Slice, reflect.Array:
			for i := cur.Len() - 1; i >= 0; i-- {
				stack.Push(cur.Index(i))
			}

		case reflect.String:
			data := []byte(cur.Interface().(string))
			_, err = buf.Write(data)

		case reflect.Bool,
			reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128:
			err = binary.Write(buf, order, cur.Interface())

		default:
			err = fmt.Errorf("Unsupported kind:" + string(kind))
		}
	}

	return buf.Bytes(), err
}
