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
