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

func (addr Address) MarshalBigEndian() (out []byte, err error) {
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

func (addr *Address) UnmarshalBigEndian(data []byte) (used int, err error) {
	size := len(data)
	if size < 4 {
		err = fmt.Errorf("Need more %d byte(s)", 4-size)
		return
	}
	addr.Type = data[0]
	switch addr.Type {
	case 1:
		if size < 7 {
			err = fmt.Errorf("Need more %d byte(s)", 7-size)
			break
		}
		addr.Addr = make([]byte, 4)
		copy(addr.Addr, data[1:5])
		addr.Port = binary.BigEndian.Uint16(data[5:7])
		used = 7
	case 3:
		length := int(data[1])
		if length == 0 {
			err = fmt.Errorf("Invalid domain name length")
			break
		} else if size < length+4 {
			err = fmt.Errorf("Need more %d byte(s)", length+4-size)
			break
		}
		addr.Addr = make([]byte, length)
		copy(addr.Addr, data[2:length+2])
		addr.Port = binary.BigEndian.Uint16(data[length+2 : length+4])
		used = length + 4
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

	if err := bin.UnmarshalBigEndian([]byte{4, 2, 0, 1, 192, 168, 1, 1, 1, 187}, &req); err != nil {
		fmt.Printf("bin.UnmarshalBigEndian error:%v", err)
	} else {
		fmt.Printf("bin.UnmarshalBigEndian Ver: %d, Cmd: %d, Target.Type: %d, Target.Addr: %v, Target.Port: %d\n",
			req.Ver, req.Cmd, req.Target.Type, req.Target.Addr, req.Target.Port)
	}
	if err := bin.UnmarshalBigEndian([]byte{5, 1, 0, 3, 15, 119, 119, 119, 46, 101, 120, 97, 109, 112, 108, 101, 46, 99, 111, 109, 4, 56}, &req); err != nil {
		fmt.Printf("bin.UnmarshalBigEndian error:%v", err)
	} else {
		fmt.Printf("bin.UnmarshalBigEndian Ver: %d, Cmd: %d, Target.Type: %d, Target.Addr: %s, Target.Port: %d\n",
			req.Ver, req.Cmd, req.Target.Type, string(req.Target.Addr), req.Target.Port)
	}

	// Output:
	// bin.MarshalBigEndian [5 1 0 1 127 0 0 1 4 56]
	// bin.MarshalBigEndian [5 1 0 3 15 119 119 119 46 101 120 97 109 112 108 101 46 99 111 109 1 187]
	// bin.UnmarshalBigEndian Ver: 4, Cmd: 2, Target.Type: 1, Target.Addr: [192 168 1 1], Target.Port: 443
	// bin.UnmarshalBigEndian Ver: 5, Cmd: 1, Target.Type: 3, Target.Addr: www.example.com, Target.Port: 1080
}
