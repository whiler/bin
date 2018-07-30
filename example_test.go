package bin_test

import (
	"bin"
	"fmt"
)

func ExampleMarshalBigEndian() {
	type IPv4 struct {
		Addr [4]byte
		Port uint16
	}

	type Greet struct {
		Ver  uint8  `bin:"1"`
		Addr *IPv4  `bin:"0"`
		Memo string `bin:"-"`
	}

	ins := Greet{
		Ver:  5,
		Addr: &IPv4{Addr: [4]byte{127, 0, 0, 1}, Port: 1080},
		Memo: "Hello,世界",
	}

	bytes, err := bin.MarshalBigEndian(ins)
	if err != nil {
		fmt.Printf("bin.MarshalBigEndian error: %v", err)
	} else {
		fmt.Println(bytes)
	}

	// Output:
	// [127 0 0 1 4 56 5]
}

func ExampleMarshalLittleEndian() {
	type IPv4 struct {
		Addr [4]byte
		Port uint16
	}

	type Greet struct {
		Ver  uint8  `bin:"1"`
		Addr *IPv4  `bin:"0"`
		Memo string `bin:"-"`
	}

	ins := Greet{
		Ver:  5,
		Addr: &IPv4{Addr: [4]byte{127, 0, 0, 1}, Port: 1080},
		Memo: "Hello,世界",
	}

	bytes, err := bin.MarshalLittleEndian(ins)
	if err != nil {
		fmt.Printf("bin.MarshalLittleEndian error: %v", err)
	} else {
		fmt.Println(bytes)
	}

	// Output:
	// [127 0 0 1 56 4 5]
}
