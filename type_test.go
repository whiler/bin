package bin

import (
	"bytes"
	"testing"
)

func TestTypesMarshal(t *testing.T) {
	for i, caze := range []struct {
		ins    interface{}
		big    []byte
		little []byte
	}{
		{Bytes8([]byte("abc")), []byte{3, 97, 98, 99}, []byte{3, 97, 98, 99}},
		{String8("abc"), []byte{3, 97, 98, 99}, []byte{3, 97, 98, 99}},
	} {
		if bs, err := MarshalBigEndian(caze.ins); err != nil {
			t.Errorf("case %d got unexcept error: %v", i, err)
		} else if !bytes.Equal(bs, caze.big) {
			t.Errorf("case %d except %v, but got %v", i, caze.big, bs)
		}
		if bs, err := MarshalLittleEndian(caze.ins); err != nil {
			t.Errorf("case %d got unexcept error: %v", i, err)
		} else if !bytes.Equal(bs, caze.little) {
			t.Errorf("case %d except %v, but got %v", i, caze.big, bs)
		}
	}
}

func TestBytes8Unmarshal(t *testing.T) {
	var bs8 Bytes8

	if err := UnmarshalBigEndian([]byte{3, 97, 98, 99}, &bs8); err != nil {
		t.Errorf("unexcept error: %v", err)
	} else if !bytes.Equal([]byte(bs8), []byte{97, 98, 99}) {
		t.Errorf("except %v, but got %v", []byte{97, 98, 99}, []byte(bs8))
	}

	if err := UnmarshalLittleEndian([]byte{4, 119, 120, 121, 122}, &bs8); err != nil {
		t.Errorf("unexcept error: %v", err)
	} else if !bytes.Equal([]byte(bs8), []byte{119, 120, 121, 122}) {
		t.Errorf("except %v, but got %v", []byte{119, 120, 121, 122}, []byte(bs8))
	}
}

func TestBytes8UnmarshalError(t *testing.T) {
	var bs8 Bytes8
	if err := UnmarshalBigEndian([]byte{3}, &bs8); err == nil {
		t.Errorf("except some error but got nil")
	}
}

func TestString8Unmarshal(t *testing.T) {
	var s8 String8

	if err := UnmarshalBigEndian([]byte{3, 97, 98, 99}, &s8); err != nil {
		t.Errorf("unexcept error: %v", err)
	} else if string(s8) != "abc" {
		t.Errorf("except %v, but got %v", "abc", string(s8))
	}

	if err := UnmarshalLittleEndian([]byte{4, 119, 120, 121, 122}, &s8); err != nil {
		t.Errorf("unexcept error: %v", err)
	} else if string(s8) != "wxyz" {
		t.Errorf("except %v, but got %v", "wxyz", string(s8))
	}
}

func TestString8UnmarshalError(t *testing.T) {
	var s8 String8
	if err := UnmarshalBigEndian([]byte{3}, &s8); err == nil {
		t.Errorf("except some error but got nil")
	}
}

func TestBytes16Marshal(t *testing.T) {
	for i, cur := range []struct {
		Instance     Bytes16
		BigEndian    []byte
		LittleEndian []byte
	}{
		{
			Bytes16{Length: 2, Value: []byte{117, 117}},
			[]byte{0, 2, 117, 117},
			[]byte{2, 0, 117, 117},
		},
	} {
		if bs, err := cur.Instance.MarshalBigEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.BigEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.BigEndian, bs)
		}
		if bs, err := cur.Instance.MarshalLittleEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.LittleEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.LittleEndian, bs)
		}
	}
}

func TestBytes16Unmarshal(t *testing.T) {
	for i, cur := range []struct {
		BigEndian    []byte
		LittleEndian []byte
		Instance     Bytes16
	}{
		{
			[]byte{0, 2, 117, 117},
			[]byte{2, 0, 117, 117},
			Bytes16{Length: 2, Value: []byte{117, 117}},
		},
	} {
		big := Bytes16{}
		if err := UnmarshalBigEndian(cur.BigEndian, &big); nil != err {
			t.Error(err)
		} else if big.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, big.Length)
		} else if !bytes.Equal(big.Value, cur.Instance.Value) {
			t.Errorf("case %d: except %v, got %v", i, cur.Instance.Value, big.Value)
		}

		little := Bytes16{}
		if err := UnmarshalLittleEndian(cur.LittleEndian, &little); nil != err {
			t.Error(err)
		} else if little.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, little.Length)
		} else if !bytes.Equal(little.Value, cur.Instance.Value) {
			t.Errorf("case %d: except %v, got %v", i, cur.Instance.Value, little.Value)
		}
	}
}

func TestBytes16UnmarshalError(t *testing.T) {
	ins := Bytes16{}
	for i, bs := range [][]byte{
		{0},
		{0, 2},
	} {
		if err := UnmarshalBigEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
		if err := UnmarshalLittleEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
	}
}

func TestBytes32Marshal(t *testing.T) {
	for i, cur := range []struct {
		Instance     Bytes32
		BigEndian    []byte
		LittleEndian []byte
	}{
		{
			Bytes32{Length: 2, Value: []byte{117, 117}},
			[]byte{0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 117, 117},
		},
	} {
		if bs, err := cur.Instance.MarshalBigEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.BigEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.BigEndian, bs)
		}
		if bs, err := cur.Instance.MarshalLittleEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.LittleEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.LittleEndian, bs)
		}
	}
}

func TestBytes32Unmarshal(t *testing.T) {
	for i, cur := range []struct {
		BigEndian    []byte
		LittleEndian []byte
		Instance     Bytes32
	}{
		{
			[]byte{0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 117, 117},
			Bytes32{Length: 2, Value: []byte{117, 117}},
		},
	} {
		big := Bytes32{}
		if err := UnmarshalBigEndian(cur.BigEndian, &big); nil != err {
			t.Error(err)
		} else if big.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, big.Length)
		} else if !bytes.Equal(big.Value, cur.Instance.Value) {
			t.Errorf("case %d: except %v, got %v", i, cur.Instance.Value, big.Value)
		}

		little := Bytes32{}
		if err := UnmarshalLittleEndian(cur.LittleEndian, &little); nil != err {
			t.Error(err)
		} else if little.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, little.Length)
		} else if !bytes.Equal(little.Value, cur.Instance.Value) {
			t.Errorf("case %d: except %v, got %v", i, cur.Instance.Value, little.Value)
		}
	}
}

func TestBytes32UnmarshalError(t *testing.T) {
	ins := Bytes32{}
	for i, bs := range [][]byte{
		{0},
		{0, 0, 0, 2},
	} {
		if err := UnmarshalBigEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
		if err := UnmarshalLittleEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
	}
}

func TestBytes64Marshal(t *testing.T) {
	for i, cur := range []struct {
		Instance     Bytes64
		BigEndian    []byte
		LittleEndian []byte
	}{
		{
			Bytes64{Length: 2, Value: []byte{117, 117}},
			[]byte{0, 0, 0, 0, 0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 0, 0, 0, 0, 117, 117},
		},
	} {
		if bs, err := cur.Instance.MarshalBigEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.BigEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.BigEndian, bs)
		}
		if bs, err := cur.Instance.MarshalLittleEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.LittleEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.LittleEndian, bs)
		}
	}
}

func TestBytes64Unmarshal(t *testing.T) {
	for i, cur := range []struct {
		BigEndian    []byte
		LittleEndian []byte
		Instance     Bytes64
	}{
		{
			[]byte{0, 0, 0, 0, 0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 0, 0, 0, 0, 117, 117},
			Bytes64{Length: 2, Value: []byte{117, 117}},
		},
	} {
		big := Bytes64{}
		if err := UnmarshalBigEndian(cur.BigEndian, &big); nil != err {
			t.Error(err)
		} else if big.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, big.Length)
		} else if !bytes.Equal(big.Value, cur.Instance.Value) {
			t.Errorf("case %d: except %v, got %v", i, cur.Instance.Value, big.Value)
		}

		little := Bytes64{}
		if err := UnmarshalLittleEndian(cur.LittleEndian, &little); nil != err {
			t.Error(err)
		} else if little.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, little.Length)
		} else if !bytes.Equal(little.Value, cur.Instance.Value) {
			t.Errorf("case %d: except %v, got %v", i, cur.Instance.Value, little.Value)
		}
	}
}

func TestBytes64UnmarshalError(t *testing.T) {
	ins := Bytes64{}
	for i, bs := range [][]byte{
		{0},
		{0, 0, 0, 0, 0, 0, 0, 2},
	} {
		if err := UnmarshalBigEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
		if err := UnmarshalLittleEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
	}
}

func TestString16Marshal(t *testing.T) {
	for i, cur := range []struct {
		Instance     String16
		BigEndian    []byte
		LittleEndian []byte
	}{
		{
			String16{Length: 2, Value: "uu"},
			[]byte{0, 2, 117, 117},
			[]byte{2, 0, 117, 117},
		},
	} {
		if bs, err := cur.Instance.MarshalBigEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.BigEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.BigEndian, bs)
		}
		if bs, err := cur.Instance.MarshalLittleEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.LittleEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.LittleEndian, bs)
		}
	}
}

func TestString16Unmarshal(t *testing.T) {
	for i, cur := range []struct {
		BigEndian    []byte
		LittleEndian []byte
		Instance     String16
	}{
		{
			[]byte{0, 2, 117, 117},
			[]byte{2, 0, 117, 117},
			String16{Length: 2, Value: "uu"},
		},
	} {
		big := String16{}
		if err := UnmarshalBigEndian(cur.BigEndian, &big); nil != err {
			t.Error(err)
		} else if big.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, big.Length)
		} else if big.Value != cur.Instance.Value {
			t.Errorf("case %d: except %s, got %s", i, cur.Instance.Value, big.Value)
		}

		little := String16{}
		if err := UnmarshalLittleEndian(cur.LittleEndian, &little); nil != err {
			t.Error(err)
		} else if little.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, little.Length)
		} else if little.Value != cur.Instance.Value {
			t.Errorf("case %d: except %s, got %s", i, cur.Instance.Value, little.Value)
		}
	}
}

func TestString16UnmarshalError(t *testing.T) {
	ins := String16{}
	for i, bs := range [][]byte{
		{0},
		{0, 2},
	} {
		if err := UnmarshalBigEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
		if err := UnmarshalLittleEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
	}
}

func TestString32Marshal(t *testing.T) {
	for i, cur := range []struct {
		Instance     String32
		BigEndian    []byte
		LittleEndian []byte
	}{
		{
			String32{Length: 2, Value: "uu"},
			[]byte{0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 117, 117},
		},
	} {
		if bs, err := cur.Instance.MarshalBigEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.BigEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.BigEndian, bs)
		}
		if bs, err := cur.Instance.MarshalLittleEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.LittleEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.LittleEndian, bs)
		}
	}
}

func TestString32Unmarshal(t *testing.T) {
	for i, cur := range []struct {
		BigEndian    []byte
		LittleEndian []byte
		Instance     String32
	}{
		{
			[]byte{0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 117, 117},
			String32{Length: 2, Value: "uu"},
		},
	} {
		big := String32{}
		if err := UnmarshalBigEndian(cur.BigEndian, &big); nil != err {
			t.Error(err)
		} else if big.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, big.Length)
		} else if big.Value != cur.Instance.Value {
			t.Errorf("case %d: except %s, got %s", i, cur.Instance.Value, big.Value)
		}

		little := String32{}
		if err := UnmarshalLittleEndian(cur.LittleEndian, &little); nil != err {
			t.Error(err)
		} else if little.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, little.Length)
		} else if little.Value != cur.Instance.Value {
			t.Errorf("case %d: except %s, got %s", i, cur.Instance.Value, little.Value)
		}
	}
}

func TestString32UnmarshalError(t *testing.T) {
	ins := String32{}
	for i, bs := range [][]byte{
		{0},
		{0, 0, 0, 2},
	} {
		if err := UnmarshalBigEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
		if err := UnmarshalLittleEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
	}
}

func TestString64Marshal(t *testing.T) {
	for i, cur := range []struct {
		Instance     String64
		BigEndian    []byte
		LittleEndian []byte
	}{
		{
			String64{Length: 2, Value: "uu"},
			[]byte{0, 0, 0, 0, 0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 0, 0, 0, 0, 117, 117},
		},
	} {
		if bs, err := cur.Instance.MarshalBigEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.BigEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.BigEndian, bs)
		}
		if bs, err := cur.Instance.MarshalLittleEndian(); nil != err {
			t.Error(err)
		} else if !bytes.Equal(bs, cur.LittleEndian) {
			t.Errorf("case %d: except %v, got %v", i, cur.LittleEndian, bs)
		}
	}
}

func TestString64Unmarshal(t *testing.T) {
	for i, cur := range []struct {
		BigEndian    []byte
		LittleEndian []byte
		Instance     String64
	}{
		{
			[]byte{0, 0, 0, 0, 0, 0, 0, 2, 117, 117},
			[]byte{2, 0, 0, 0, 0, 0, 0, 0, 117, 117},
			String64{Length: 2, Value: "uu"},
		},
	} {
		big := String64{}
		if err := UnmarshalBigEndian(cur.BigEndian, &big); nil != err {
			t.Error(err)
		} else if big.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, big.Length)
		} else if big.Value != cur.Instance.Value {
			t.Errorf("case %d: except %s, got %s", i, cur.Instance.Value, big.Value)
		}

		little := String64{}
		if err := UnmarshalLittleEndian(cur.LittleEndian, &little); nil != err {
			t.Error(err)
		} else if little.Length != cur.Instance.Length {
			t.Errorf("case %d: except %d, got %d", i, cur.Instance.Length, little.Length)
		} else if little.Value != cur.Instance.Value {
			t.Errorf("case %d: except %s, got %s", i, cur.Instance.Value, little.Value)
		}
	}
}

func TestString64UnmarshalError(t *testing.T) {
	ins := String64{}
	for i, bs := range [][]byte{
		{0},
		{0, 0, 0, 0, 0, 0, 0, 2},
	} {
		if err := UnmarshalBigEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
		if err := UnmarshalLittleEndian(bs, &ins); nil == err {
			t.Errorf("case %d, except some error", i)
		}
	}
}
