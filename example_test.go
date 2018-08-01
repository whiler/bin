package bin_test

import (
	"bin"
	"encoding/binary"
	"fmt"
)

type Address struct {
	Type byte
	Addr []byte
	Port uint16
}

func (addr *Address) MarshalBigEndian() (out []byte, err error) {
	switch addr.Type {
	case 1:
		out = make([]byte, 7)
		out[0] = addr.Type
		copy(out[1:5], addr.Addr)
		binary.BigEndian.PutUint16(out[5:7], addr.Port)
	case 3:
		length := len(addr.Addr)
		out = make([]byte, length+4)
		out[0] = addr.Type
		out[1] = byte(length)
		copy(out[2:length+2], addr.Addr)
		binary.BigEndian.PutUint16(out[length+2:length+4], addr.Port)
	default:
		err = fmt.Errorf("Invalid Type %d", addr.Type)
	}
	return
}

type Request struct {
	Ver    byte
	Cmd    byte
	Rsv    byte
	Target *Address
	Memo   string `bin:"-"`
}

func Example() {
	req := Request{
		Ver:  5,
		Cmd:  1,
		Memo: "this is memo",
	}

	req.Target = &Address{
		Type: 1,
		Addr: []byte{127, 0, 0, 1},
		Port: 1080,
	}
	if bs, err := bin.MarshalBigEndian(req); err != nil {
		fmt.Printf("bin.MarshalBigEndian error:%v", err)
	} else {
		fmt.Println("bin.MarshalBigEndian", bs)
	}

	req.Target = &Address{
		Type: 3,
		Addr: []byte("www.example.com"),
		Port: 443,
	}
	if bs, err := bin.MarshalBigEndian(req); err != nil {
		fmt.Printf("bin.MarshalBigEndian error:%v", err)
	} else {
		fmt.Println("bin.MarshalBigEndian", bs)
	}

	// Output:
	// bin.MarshalBigEndian [5 1 0 1 127 0 0 1 4 56]
	// bin.MarshalBigEndian [5 1 0 3 15 119 119 119 46 101 120 97 109 112 108 101 46 99 111 109 1 187]
}
