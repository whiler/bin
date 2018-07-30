package bin_test

import (
	"bin"
	"fmt"
)

type BigAuth struct {
	Count   int
	Methods []uint8
}

func (auth BigAuth) MarshalBigEndian() ([]byte, error) {
	data := make([]byte, auth.Count+1)
	data[0] = byte(auth.Count)
	for i := 0; i < auth.Count; i++ {
		data[i+1] = byte(31)
	}
	return data, nil
}

func Example_customMarshalBigEndian() {
	ins := BigAuth{Count: 3, Methods: []uint8{2, 5, 7}}
	bs, err := bin.MarshalBigEndian(ins)
	if err != nil {
		fmt.Printf("bin.MarshalBigEndian error:%v", err)
	} else {
		fmt.Println(bs)
	}

	// Output:
	// [3 31 31 31]
}
