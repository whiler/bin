package bin_test

import (
	"bin"
	"errors"
	"fmt"
)

type UnBigAuth struct {
	Count   int
	Methods []uint8
}

func (auth *UnBigAuth) UnmarshalBigEndian(data []byte) (used int, err error) {
	size := len(data)
	if size == 0 {
		err = errors.New("Need more bytes")
		return
	}
	auth.Count = int(data[0])
	if size < auth.Count+1 {
		err = errors.New("Need more bytes")
		return
	}
	auth.Methods = make([]uint8, auth.Count)
	for i := 0; i < auth.Count; i++ {
		auth.Methods[i] = data[i+1]
	}
	used = auth.Count + 1
	return
}

func Example_customUnmarshalBigEndian() {
	bytes := []byte{3, 2, 5, 7}
	ins := UnBigAuth{}
	err := bin.UnmarshalBigEndian(bytes, &ins)
	if err != nil {
		fmt.Printf("bin.UnmarshalBigEndian error:%v", err)
	} else {
		fmt.Println(ins)
	}

	// Output:
	// {3 [2 5 7]}
}
