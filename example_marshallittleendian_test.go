package bin_test

import (
	"bin"
	"fmt"
)

type LittleAuth struct {
	Count   int
	Methods []uint8
}

func (auth LittleAuth) MarshalLittleEndian() ([]byte, error) {
	data := make([]byte, auth.Count+1)
	data[0] = byte(auth.Count)
	for i := 0; i < auth.Count; i++ {
		data[i+1] = byte(13)
	}
	return data, nil
}

func Example_customMarshalLittleEndian() {
	ins := LittleAuth{Count: 3, Methods: []uint8{2, 5, 7}}
	bs, err := bin.MarshalLittleEndian(ins)
	if err != nil {
		fmt.Printf("bin.MarshalLittleEndian error:%v", err)
	} else {
		fmt.Println(bs)
	}

	// Output:
	// [3 13 13 13]
}
