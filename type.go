package bin

import (
	"fmt"
)

// Bytes8 defines a common byte slice type, which max length is 255.
//
// +-----+-------+--------+   +--------+
// | LEN |item-0 | item-1 |...| item-n |
// +-----+-------+--------+   +--------+
type Bytes8 []byte

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (bs8 Bytes8) MarshalBigEndian() ([]byte, error) {
	return append([]byte{byte(len(bs8))}, []byte(bs8)...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (bs8 Bytes8) MarshalLittleEndian() ([]byte, error) {
	return append([]byte{byte(len(bs8))}, []byte(bs8)...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (bs8 *Bytes8) UnmarshalBigEndian(data []byte) (used int, err error) {
	size := len(data)
	length := int(data[0])
	if size < length+1 {
		err = fmt.Errorf("Need more %d byte(s)", length+1-size)
		return
	}
	*bs8 = make([]byte, length)
	copy(*bs8, data[1:length+1])
	used = length + 1
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (bs8 *Bytes8) UnmarshalLittleEndian(data []byte) (used int, err error) {
	return bs8.UnmarshalBigEndian(data)
}

// String8 defines a common string type, which max length is 255.
//
// +-----+-------+--------+   +--------+
// | LEN |char-0 | char-1 |...| char-n |
// +-----+-------+--------+   +--------+
type String8 string

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (s8 String8) MarshalBigEndian() ([]byte, error) {
	return append([]byte{byte(len(s8))}, []byte(s8)...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (s8 String8) MarshalLittleEndian() ([]byte, error) {
	return append([]byte{byte(len(s8))}, []byte(s8)...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (s8 *String8) UnmarshalBigEndian(data []byte) (used int, err error) {
	size := len(data)
	length := int(data[0])
	if size < length+1 {
		err = fmt.Errorf("Need more %d byte(s)", length+1-size)
		return
	}
	*s8 = String8(data[1 : length+1])
	used = length + 1
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (s8 *String8) UnmarshalLittleEndian(data []byte) (used int, err error) {
	return s8.UnmarshalBigEndian(data)
}
