package bin

import (
	"bytes"
	"fmt"
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

func TestMarshalLittleEndian(t *testing.T) {
	for i, caze := range []struct {
		ins    interface{}
		except []byte
	}{
		{int16(255), []byte{255, 0}},
		{[]int16{255, 255}, []byte{255, 0, 255, 0}},
		{[2]int16{255, 255}, []byte{255, 0, 255, 0}},
		{struct{ I16 int16 }{255}, []byte{255, 0}},
		{struct{ I16s []int16 }{[]int16{255, 255}}, []byte{255, 0, 255, 0}},
		{struct{ I16a [2]int16 }{[2]int16{255, 255}}, []byte{255, 0, 255, 0}},
		{struct{ I16p *[]int16 }{&[]int16{255, 255}}, []byte{255, 0, 255, 0}},
	} {
		if bs, e := MarshalLittleEndian(caze.ins); e != nil {
			t.Errorf("case %d got unexcepted error %v", i, e)
		} else if !bytes.Equal(bs, caze.except) {
			t.Errorf("case %d except %v but got %v", i, caze.except, bs)
		}
	}
}

func TestMarshalBigEndian(t *testing.T) {
	for i, caze := range []struct {
		ins    interface{}
		except []byte
	}{
		{int16(255), []byte{0, 255}},
		{[]int16{255, 255}, []byte{0, 255, 0, 255}},
		{[2]int16{255, 255}, []byte{0, 255, 0, 255}},
		{struct{ I16 int16 }{255}, []byte{0, 255}},
		{struct{ I16s []int16 }{[]int16{255, 255}}, []byte{0, 255, 0, 255}},
		{struct{ I16a [2]int16 }{[2]int16{255, 255}}, []byte{0, 255, 0, 255}},
		{struct{ I16p *[]int16 }{&[]int16{255, 255}}, []byte{0, 255, 0, 255}},
	} {
		if bs, e := MarshalBigEndian(caze.ins); e != nil {
			t.Errorf("case %d got unexcepted error %v", i, e)
		} else if !bytes.Equal(bs, caze.except) {
			t.Errorf("case %d except %v but got %v", i, caze.except, bs)
		}
	}
}

type okMarshaler struct{}

func (ok okMarshaler) MarshalBigEndian() ([]byte, error) {
	return []byte("ok"), nil
}

func (ok okMarshaler) MarshalLittleEndian() ([]byte, error) {
	return []byte("ko"), nil
}

type koMarshaler struct{}

func (ko koMarshaler) MarshalBigEndian() ([]byte, error) {
	return []byte{}, fmt.Errorf("Test Error")
}

func (ko koMarshaler) MarshalLittleEndian() ([]byte, error) {
	return []byte{}, fmt.Errorf("Test Error")
}

func TestMarshaler(t *testing.T) {
	for _, ok := range []interface{}{
		okMarshaler{},
		&okMarshaler{},
		struct{ I *okMarshaler }{},
		&struct{ I *okMarshaler }{},
	} {
		if bs, e := MarshalBigEndian(ok); e != nil {
			t.Errorf("got unexcepted error %v", e)
		} else if !bytes.Equal(bs, []byte("ok")) {
			t.Errorf("except %v but got %v", []byte("ok"), bs)
		}
		if bs, e := MarshalLittleEndian(ok); e != nil {
			t.Errorf("got unexcepted error %v", e)
		} else if !bytes.Equal(bs, []byte("ko")) {
			t.Errorf("except %v but got %v", []byte("ko"), bs)
		}
	}
}

func TestMarshalerError(t *testing.T) {
	for _, ko := range []interface{}{
		koMarshaler{},
		&koMarshaler{},
		struct{ I *koMarshaler }{},
		&struct{ I *koMarshaler }{},
	} {
		if _, e := MarshalBigEndian(ko); e == nil {
			t.Errorf("except error but got nil")
		}
		if _, e := MarshalLittleEndian(ko); e == nil {
			t.Errorf("except error but got nil")
		}
	}
}

func TestMarshalTo(t *testing.T) {
	var (
		ins          uint16 = 1080
		exceptBig           = []byte{4, 56}
		exceptLittle        = []byte{56, 4}
		buffer              = new(bytes.Buffer)
	)
	if err := MarshalBigEndianTo(buffer, ins); err != nil {
		t.Errorf("unexcepted error: %v", err)
	} else if bs := buffer.Bytes(); !bytes.Equal(bs, exceptBig) {
		t.Errorf("except %v, but got %v,", exceptBig, bs)
	}

	buffer.Reset()

	if err := MarshalLittleEndianTo(buffer, ins); err != nil {
		t.Errorf("unexcepted error: %v", err)
	} else if bs := buffer.Bytes(); !bytes.Equal(bs, exceptLittle) {
		t.Errorf("except %v, but got %v,", exceptLittle, bs)
	}
}
