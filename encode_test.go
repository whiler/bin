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
