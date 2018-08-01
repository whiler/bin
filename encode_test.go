package bin

import (
	"bytes"
	"testing"
)

func TestMarshalTypes(t *testing.T) {
	type pair struct {
		Dst, Src uint16
	}
	type ipv4 struct {
		Addr [4]byte
		Port uint16
	}
	type inTest struct {
		ID   uint8
		Addr *ipv4
	}

	for i, caze := range []struct {
		ins    interface{}
		except []byte
	}{
		{true, []byte{1}},
		{false, []byte{0}},
		{int8(127), []byte{127}},
		{int16(255), []byte{0, 255}},
		{int32(65535), []byte{0, 0, 255, 255}},
		{int64(4294967295), []byte{0, 0, 0, 0, 255, 255, 255, 255}},
		{float32(3.1415926535897932384626433), []byte{64, 73, 15, 219}},
		{float64(3.1415926535897932384626433), []byte{64, 9, 33, 251, 84, 68, 45, 24}},
		{"string", []byte("string")},
		{[]uint16{443, 1080}, []byte{1, 187, 4, 56}},
		{[2]uint16{443, 1080}, []byte{1, 187, 4, 56}},
		{pair{443, 1080}, []byte{1, 187, 4, 56}},
		{&pair{443, 1080}, []byte{1, 187, 4, 56}},
		{inTest{ID: 47}, []byte{47, 0, 0, 0, 0, 0, 0}},
		{inTest{ID: 47, Addr: &ipv4{Addr: [4]byte{127, 0, 0, 1}}}, []byte{47, 127, 0, 0, 1, 0, 0}},
		{inTest{ID: 47, Addr: &ipv4{Port: 1080}}, []byte{47, 0, 0, 0, 0, 4, 56}},
	} {
		if bs, e := MarshalBigEndian(caze.ins); e != nil {
			t.Errorf("case %d got unexcepted error %v", i, e)
		} else if !bytes.Equal(bs, caze.except) {
			t.Errorf("case %d except %v but got %v", i, caze.except, bs)
		}
	}
}

func TestUnsupportedMarshalTypes(t *testing.T) {
	for i, caze := range []interface{}{
		nil,
		int(-1024),
		uint(1024),
	} {
		if _, e := MarshalBigEndian(caze); e == nil {
			t.Errorf("case %d excepted some error but got nil", i)
		}
	}
}

func TestMarshalTags(t *testing.T) {
	for i, caze := range []struct {
		ins    interface{}
		except []byte
		err    bool
	}{
		{struct {
			First  byte
			Second byte
			Third  byte
		}{1, 2, 3}, []byte{1, 2, 3}, false},
		{struct {
			First  byte `bin:"2"`
			Second byte `bin:"1"`
			Third  byte `bin:"0"`
		}{1, 2, 3}, []byte{3, 2, 1}, false},
		{struct {
			First  byte `bin:"-"`
			Second byte `bin:"1"`
			Third  byte `bin:"0"`
		}{1, 2, 3}, []byte{3, 2}, false},
		{struct {
			First  byte `bin:"2"`
			Second byte `bin:"-"`
			Third  byte `bin:"0"`
		}{1, 2, 3}, []byte{}, true},
		{struct {
			First  byte `bin:"3"`
			Second byte `bin:"1"`
			Third  byte `bin:"0"`
		}{1, 2, 3}, []byte{}, true},
		{struct {
			First  byte `bin:"x"`
			Second byte `bin:"1"`
			Third  byte `bin:"0"`
		}{1, 2, 3}, []byte{}, true},
		{struct {
			First  byte `bin:"1"`
			Second byte `bin:"1"`
			Third  byte `bin:"0"`
		}{1, 2, 3}, []byte{}, true},
	} {
		if bs, e := MarshalBigEndian(caze.ins); !caze.err && e != nil {
			t.Errorf("case %d unexcepted error %v", i, e)
		} else if caze.err && e == nil {
			t.Errorf("case %d excepted some error but got nil", i)
		} else if !bytes.Equal(bs, caze.except) {
			t.Errorf("case %d except %v but got %v", i, caze.except, bs)
		}
	}
}
