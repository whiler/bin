package bin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalBool(t *testing.T) {
	var (
		ins    bool
		except bool = true
	)
	if e := UnmarshalBigEndian([]byte{1}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalInt8(t *testing.T) {
	var (
		ins    int8
		except int8 = 64
	)
	if e := UnmarshalBigEndian([]byte{64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalInt16(t *testing.T) {
	var (
		ins    int16
		except int16 = 64
	)
	if e := UnmarshalBigEndian([]byte{0, 64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalInt32(t *testing.T) {
	var (
		ins    int32
		except int32 = 64
	)
	if e := UnmarshalBigEndian([]byte{0, 0, 0, 64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalInt64(t *testing.T) {
	var (
		ins    int64
		except int64 = 64
	)
	if e := UnmarshalBigEndian([]byte{0, 0, 0, 0, 0, 0, 0, 64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalUint8(t *testing.T) {
	var (
		ins    uint8
		except uint8 = 64
	)
	if e := UnmarshalBigEndian([]byte{64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalUint16(t *testing.T) {
	var (
		ins    uint16
		except uint16 = 64
	)
	if e := UnmarshalBigEndian([]byte{0, 64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalUint32(t *testing.T) {
	var (
		ins    uint32
		except uint32 = 64
	)
	if e := UnmarshalBigEndian([]byte{0, 0, 0, 64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalUint64(t *testing.T) {
	var (
		ins    uint64
		except uint64 = 64
	)
	if e := UnmarshalBigEndian([]byte{0, 0, 0, 0, 0, 0, 0, 64}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalFloat32(t *testing.T) {
	var (
		ins    float32
		except float32 = 64
	)
	if e := UnmarshalBigEndian([]byte{66, 128, 0, 0}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalFloat64(t *testing.T) {
	var (
		ins    float64
		except float64 = 64
	)
	if e := UnmarshalBigEndian([]byte{64, 80, 0, 0, 0, 0, 0, 0}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalComplex64(t *testing.T) {
	var (
		ins    complex64
		except complex64 = complex(float32(32), float32(32))
	)
	if e := UnmarshalBigEndian([]byte{66, 0, 0, 0, 66, 0, 0, 0}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalComplex128(t *testing.T) {
	var (
		ins    complex128
		except complex128 = complex(float64(64), float64(64))
	)
	if e := UnmarshalBigEndian([]byte{64, 80, 0, 0, 0, 0, 0, 0, 64, 80, 0, 0, 0, 0, 0, 0}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if ins != except {
		t.Errorf("except %v, but got %v", except, ins)
	}
}

func TestUnmarshalStruct(t *testing.T) {
	type inTest struct {
		ID   uint32
		Port uint16
	}
	var (
		ins    inTest
		except inTest = inTest{2147482859, 1080}
	)
	if e := UnmarshalBigEndian([]byte{127, 255, 252, 235, 4, 56}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if !reflect.DeepEqual(ins, except) {
		t.Errorf("except %#v, but got %#v", except, ins)
	}
}

func TestUnmarshalSlice(t *testing.T) {
	var (
		ins    []byte = make([]byte, 6)
		except []byte = []byte("decode")
	)
	if e := UnmarshalBigEndian([]byte("decode"), &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if !reflect.DeepEqual(ins, except) {
		t.Errorf("except %#v, but got %#v", except, ins)
	}
}

func TestUnmarshalArray(t *testing.T) {
	var (
		ins    [4]byte
		except [4]byte = [4]byte{127, 255, 252, 235}
	)
	if e := UnmarshalBigEndian([]byte{127, 255, 252, 235}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if !reflect.DeepEqual(ins, except) {
		t.Errorf("except %#v, but got %#v", except, ins)
	}
}

func TestUnmarshalPtr(t *testing.T) {
	type pair struct {
		Key   uint16
		Value uint16
	}
	type inTest struct {
		ID   *pair
		Port uint16
	}
	var (
		ins    inTest
		except inTest = inTest{&pair{32767, 64747}, 1080}
	)
	if e := UnmarshalBigEndian([]byte{127, 255, 252, 235, 4, 56}, &ins); e != nil {
		t.Errorf("unexcept error: %v", e)
	} else if !reflect.DeepEqual(ins, except) {
		t.Errorf("except %#v, but got %#v", except, ins)
	}
}

func TestUnmarshalUnsupportedKind(t *testing.T) {
	var (
		ins string
	)
	if e := UnmarshalBigEndian([]byte("Unsupported"), &ins); e == nil {
		t.Errorf("except some error, but got nil")
	}
}

func TestUnmarshalNeedMoreBytes(t *testing.T) {
	var (
		ins uint32
	)
	if e := UnmarshalBigEndian([]byte{0, 0, 0}, &ins); e == nil {
		t.Errorf("except some error, but got nil")
	}
}

func TestUnmarshalInvalidType(t *testing.T) {
	var (
		ins *uint32
	)
	if e := UnmarshalBigEndian([]byte{0}, ins); e == nil {
		t.Errorf("except some error, but got nil")
	}
}

type okUnmarshaler struct{ B byte }

func (ok *okUnmarshaler) UnmarshalBigEndian(data []byte) (used int, err error) {
	if len(data) > 0 {
		ok.B = data[0]
		used = 1
	} else {
		err = fmt.Errorf("Need more 1 byte")
	}
	return
}

func (ok *okUnmarshaler) UnmarshalLittleEndian(data []byte) (used int, err error) {
	if len(data) > 0 {
		ok.B = data[0]
		used = 1
	} else {
		err = fmt.Errorf("Need more 1 byte")
	}
	return
}

func TestUnmarshaler(t *testing.T) {
	ins := okUnmarshaler{}

	if err := UnmarshalBigEndian([]byte{71}, &ins); err != nil {
		t.Errorf("unexcept error: %v", err)
	} else if ins.B != 71 {
		t.Errorf("except %d, but got %d", 71, ins.B)
	}

	if err := UnmarshalLittleEndian([]byte{17}, &ins); err != nil {
		t.Errorf("unexcept error: %v", err)
	} else if ins.B != 17 {
		t.Errorf("except %d, but got %d", 17, ins.B)
	}
}

type koUnmarshaler struct{}

func (ko *koUnmarshaler) UnmarshalBigEndian(data []byte) (used int, err error) {
	err = fmt.Errorf("Test Error")
	return
}

func (ko *koUnmarshaler) UnmarshalLittleEndian(data []byte) (used int, err error) {
	err = fmt.Errorf("Test Error")
	return
}

func TestUnmarshalerError(t *testing.T) {
	ins := koUnmarshaler{}

	if err := UnmarshalBigEndian([]byte{}, &ins); err == nil {
		t.Errorf("except some error but got nil")
	}

	if err := UnmarshalLittleEndian([]byte{}, &ins); err == nil {
		t.Errorf("except some error but got nil")
	}
}
