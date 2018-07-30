package bin

import (
	"bytes"
	"errors"
	"testing"
)

func TestSupportedMarshalTypes(t *testing.T) {
	type Tpe struct {
		ID   int8
		Name string
	}
	type Tpes struct {
		Bool       bool
		Int8       int8
		Int16      int16
		Int32      int32
		Int64      int64
		Uint8      uint8
		Uint16     uint16
		Uint32     uint32
		Uint64     uint64
		Float32    float32
		Float64    float64
		Complex64  complex64
		Complex128 complex128
		Array      [4]Tpe
		Slice      []Tpe
		Ptr        *Tpe
		String     string
		Struct     Tpe
	}

	ins := Tpes{}
	except := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if bs, err := MarshalBigEndian(ins); err != nil {
		t.Error(err)
	} else if !bytes.Equal(bs, except) {
		t.Errorf("except: %v\ngot: %v", except, bs)
	}

	if bs, err := MarshalLittleEndian(ins); err != nil {
		t.Error(err)
	} else if !bytes.Equal(bs, except) {
		t.Errorf("except: %v\ngot: %v", except, bs)
	}
}

type testMarshalerError uint8

func (i testMarshalerError) MarshalBigEndian() ([]byte, error) {
	return nil, errors.New("MarshalBigEndian Error")
}

func (i testMarshalerError) MarshalLittleEndian() ([]byte, error) {
	return nil, errors.New("MarshalLittleEndian Error")
}

func TestBigEndianMarshalerError(t *testing.T) {
	var ins testMarshalerError
	_, err := MarshalBigEndian(ins)
	if err == nil {
		t.Error("except MarshalBigEndian Error, got nil")
	}
}

func TestLittleEndianMarshalerError(t *testing.T) {
	var ins testMarshalerError
	_, err := MarshalLittleEndian(ins)
	if err == nil {
		t.Error("except MarshalLittleEndian Error, got nil")
	}
}

func TestMarshalBigEndianUnsupportedError(t *testing.T) {
	var (
		i int
		u uint
		m map[int]string = make(map[int]string)
	)

	if _, e := MarshalBigEndian(i); e == nil {
		t.Error("except Unsupported kind error, got nil")
	}
	if _, e := MarshalBigEndian(u); e == nil {
		t.Error("except Unsupported kind error, got nil")
	}
	if _, e := MarshalBigEndian(m); e == nil {
		t.Error("except Unsupported kind error, got nil")
	}
}

func TestMarshalBigEndianInvalidTagError(t *testing.T) {
	type syntaxTag struct {
		I byte `bin:"ko"`
	}
	syntax := syntaxTag{}
	if _, e := MarshalBigEndian(syntax); e == nil {
		t.Error("except syntax error, got nil")
	}

	type overTag struct {
		I byte `bin:"2"`
	}
	over := overTag{}
	if _, e := MarshalBigEndian(over); e == nil {
		t.Error("except Field index out of range error, got nil")
	}

	type dupTag struct {
		I  byte `bin:"1"`
		II byte `bin:"1"`
	}
	dup := dupTag{}
	if _, e := MarshalBigEndian(dup); e == nil {
		t.Error("except Field index duplicated error, got nil")
	}

	type emptyTag struct {
		I   byte `bin:"0"`
		II  byte `bin:"-"`
		III byte `bin:"2"`
	}
	empty := emptyTag{}
	if _, e := MarshalBigEndian(empty); e == nil {
		t.Error("except Field index is invalid error, got nil")
	}
}

func TestMarshalBigEndianInvalidValue(t *testing.T) {
	if _, e := MarshalBigEndian(nil); e == nil {
		t.Error("except Unexcepted error, got nil")
	}
}
