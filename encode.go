package bin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

// BigEndianMarshaler is the interface implemented by types that can marshal themselves into valid big-endian binary data.
type BigEndianMarshaler interface {
	MarshalBigEndian() ([]byte, error)
}

// LittleEndianMarshaler is the interface implemented by types that can marshal themselves into valid little-endian binary data.
type LittleEndianMarshaler interface {
	MarshalLittleEndian() ([]byte, error)
}

// MarshalBigEndian returns the big-endian encoding binary data of ins.
//
// MarshalBigEndian traverses the value ins recursively.
// If an encountered value implements the BigEndianMarshaler interface,
// MarshalBigEndian calls its MarshalBigEndian method to produce big-endian binary data.
func MarshalBigEndian(ins interface{}) ([]byte, error) {
	var buffer = new(bytes.Buffer)
	if err := marshal(buffer, ins, binary.BigEndian, bigEndianMarshalerType, bigEndianMarshaler); err != nil {
		return []byte{}, err
	} else {
		return buffer.Bytes(), err
	}
}

// MarshalLittleEndian returns the little-endian encoding binary data of ins.
//
// MarshalLittleEndian traverses the value ins recursively.
// If an encountered value implements the LittleEndianMarshaler interface,
// MarshalLittleEndian calls its MarshalLittleEndian method to produce little-endian binary data.
func MarshalLittleEndian(ins interface{}) ([]byte, error) {
	var buffer = new(bytes.Buffer)
	if err := marshal(buffer, ins, binary.LittleEndian, littleEndianMarshalerType, littleEndianMarshaler); err != nil {
		return []byte{}, err
	} else {
		return buffer.Bytes(), err
	}
}

// MarshalBigEndianTo writes the big-endian encoding binary data of ins into writer.
//
// MarshalBigEndianTo traverses the value ins recursively.
// If an encountered value implements the BigEndianMarshaler interface,
// MarshalBigEndianTo calls its MarshalBigEndian method to produce big-endian binary data.
func MarshalBigEndianTo(writer io.Writer, ins interface{}) error {
	return marshal(writer, ins, binary.BigEndian, bigEndianMarshalerType, bigEndianMarshaler)
}

// MarshalLittleEndianTo writes the little-endian encoding binary data of ins into writer.
//
// MarshalLittleEndianTo traverses the value ins recursively.
// If an encountered value implements the LittleEndianMarshaler interface,
// MarshalLittleEndianTo calls its MarshalLittleEndian method to produce little-endian binary data.
func MarshalLittleEndianTo(writer io.Writer, ins interface{}) error {
	return marshal(writer, ins, binary.LittleEndian, littleEndianMarshalerType, littleEndianMarshaler)
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

func marshal(writer io.Writer, ins interface{}, order binary.ByteOrder, marshalerType reflect.Type, marshaler marshalerFunc) error {
	var (
		stack valueStack
		cur   reflect.Value
		tpe   reflect.Type
		kind  reflect.Kind
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
		kind = cur.Kind()
		if tpe.Implements(marshalerType) {
			if kind == reflect.Ptr && cur.IsNil() {
				cur = reflect.New(tpe.Elem())
			}
			if data, e := marshaler(cur.Interface()); e != nil {
				err = e
			} else {
				_, err = writer.Write(data)
			}
			continue
		}

		switch kind {
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
			_, err = writer.Write([]byte(cur.String()))

		case reflect.Bool,
			reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128:
			err = binary.Write(writer, order, cur.Interface())

		default:
			err = fmt.Errorf("Unsupported kind %s", kind)
		}
	}

	return err
}
