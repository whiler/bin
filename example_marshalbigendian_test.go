package bin_test

import (
	"bin"
	"fmt"
)

type Auth struct {
	Count   int
	Methods []uint8
}

func (auth Auth) MarshalBigEndian() ([]byte, error) {
	data := make([]byte, auth.Count+1)
	data[0] = byte(auth.Count)
	for i := 0; i < auth.Count; i++ {
		data[i+1] = byte(11)
	}
	return data, nil
}

func Example_customMarshalBigEndian() {
	ins := Auth{Count: 3, Methods: []uint8{2, 5, 7}}
	bs, err := bin.MarshalBigEndian(ins)
	if err != nil {
		fmt.Printf("bin.MarshalBigEndian error:%v", err)
	} else {
		fmt.Println(bs)
	}

	// Output:
	// [3 11 11 11]
}
