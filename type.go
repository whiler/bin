package bin

import (
	"encoding/binary"
	"fmt"
)

// Bytes8 defines a common byte slice type, which max length is 255.
//
//	+-----+-------+--------+   +--------+
// 	| LEN |item-0 | item-1 |...| item-n |
// 	+-----+-------+--------+   +--------+
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
//	+-----+-------+--------+   +--------+
// 	| LEN |char-0 | char-1 |...| char-n |
// 	+-----+-------+--------+   +--------+
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

// Bytes16 defines a common byte slice type, the max length is math.MaxUint16.
//
// +--....--+-------+--------+   +--------+
// | length |byte-0 | byte-1 |...| byte-n |
// +--....--+-------+--------+   +--------+
type Bytes16 struct {
	Length uint16
	Value  []byte
}

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (bs16 Bytes16) MarshalBigEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(bs16.Length))
	binary.BigEndian.PutUint16(mem, bs16.Length)
	return append(mem, bs16.Value...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (bs16 Bytes16) MarshalLittleEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(bs16.Length))
	binary.LittleEndian.PutUint16(mem, bs16.Length)
	return append(mem, bs16.Value...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (bs16 *Bytes16) UnmarshalBigEndian(data []byte) (used int, err error) {
	var size int = binary.Size(bs16.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	bs16.Length = binary.BigEndian.Uint16(data[:size])
	if uint16(length-size) < bs16.Length {
		err = fmt.Errorf("Need more %d byte(s)", bs16.Length+uint16(size)-uint16(length))
		return
	}
	bs16.Value = make([]byte, int(bs16.Length))
	used = size + int(bs16.Length)
	copy(bs16.Value, data[size:used])
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (bs16 *Bytes16) UnmarshalLittleEndian(data []byte) (used int, err error) {
	var size int = binary.Size(bs16.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	bs16.Length = binary.LittleEndian.Uint16(data[:size])
	if uint16(length-size) < bs16.Length {
		err = fmt.Errorf("Need more %d byte(s)", bs16.Length+uint16(size)-uint16(length))
		return
	}
	bs16.Value = make([]byte, int(bs16.Length))
	used = size + int(bs16.Length)
	copy(bs16.Value, data[size:used])
	return
}

// Bytes32 defines a common byte slice type, the max length is math.MaxUint32.
//
// +--....--+-------+--------+   +--------+
// | length |byte-0 | byte-1 |...| byte-n |
// +--....--+-------+--------+   +--------+
type Bytes32 struct {
	Length uint32
	Value  []byte
}

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (bs32 Bytes32) MarshalBigEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(bs32.Length))
	binary.BigEndian.PutUint32(mem, bs32.Length)
	return append(mem, bs32.Value...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (bs32 Bytes32) MarshalLittleEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(bs32.Length))
	binary.LittleEndian.PutUint32(mem, bs32.Length)
	return append(mem, bs32.Value...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (bs32 *Bytes32) UnmarshalBigEndian(data []byte) (used int, err error) {
	var size int = binary.Size(bs32.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	bs32.Length = binary.BigEndian.Uint32(data[:size])
	if uint32(length-size) < bs32.Length {
		err = fmt.Errorf("Need more %d byte(s)", bs32.Length+uint32(size)-uint32(length))
		return
	}
	bs32.Value = make([]byte, int(bs32.Length))
	used = size + int(bs32.Length)
	copy(bs32.Value, data[size:used])
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (bs32 *Bytes32) UnmarshalLittleEndian(data []byte) (used int, err error) {
	var size int = binary.Size(bs32.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	bs32.Length = binary.LittleEndian.Uint32(data[:size])
	if uint32(length-size) < bs32.Length {
		err = fmt.Errorf("Need more %d byte(s)", bs32.Length+uint32(size)-uint32(length))
		return
	}
	bs32.Value = make([]byte, int(bs32.Length))
	used = size + int(bs32.Length)
	copy(bs32.Value, data[size:used])
	return
}

// Bytes64 defines a common byte slice type, the max length is math.MaxUint64.
//
// +--....--+-------+--------+   +--------+
// | length |byte-0 | byte-1 |...| byte-n |
// +--....--+-------+--------+   +--------+
type Bytes64 struct {
	Length uint64
	Value  []byte
}

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (bs64 Bytes64) MarshalBigEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(bs64.Length))
	binary.BigEndian.PutUint64(mem, bs64.Length)
	return append(mem, bs64.Value...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (bs64 Bytes64) MarshalLittleEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(bs64.Length))
	binary.LittleEndian.PutUint64(mem, bs64.Length)
	return append(mem, bs64.Value...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (bs64 *Bytes64) UnmarshalBigEndian(data []byte) (used int, err error) {
	var size int = binary.Size(bs64.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	bs64.Length = binary.BigEndian.Uint64(data[:size])
	if uint64(length-size) < bs64.Length {
		err = fmt.Errorf("Need more %d byte(s)", bs64.Length+uint64(size)-uint64(length))
		return
	}
	bs64.Value = make([]byte, int(bs64.Length))
	used = size + int(bs64.Length)
	copy(bs64.Value, data[size:used])
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (bs64 *Bytes64) UnmarshalLittleEndian(data []byte) (used int, err error) {
	var size int = binary.Size(bs64.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	bs64.Length = binary.LittleEndian.Uint64(data[:size])
	if uint64(length-size) < bs64.Length {
		err = fmt.Errorf("Need more %d byte(s)", bs64.Length+uint64(size)-uint64(length))
		return
	}
	bs64.Value = make([]byte, int(bs64.Length))
	used = size + int(bs64.Length)
	copy(bs64.Value, data[size:used])
	return
}

// String16 defines a common string type, which max length is math.MaxUint16.
//
// +--....--+-------+--------+   +--------+
// | length |char-0 | char-1 |...| char-n |
// +--....--+-------+--------+   +--------+
type String16 struct {
	Length uint16
	Value  string
}

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (s16 String16) MarshalBigEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(s16.Length))
	binary.BigEndian.PutUint16(mem, s16.Length)
	return append(mem, []byte(s16.Value)...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (s16 String16) MarshalLittleEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(s16.Length))
	binary.LittleEndian.PutUint16(mem, s16.Length)
	return append(mem, []byte(s16.Value)...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (s16 *String16) UnmarshalBigEndian(data []byte) (used int, err error) {
	var size int = binary.Size(s16.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	s16.Length = binary.BigEndian.Uint16(data[:size])
	if uint16(length-size) < s16.Length {
		err = fmt.Errorf("Need more %d byte(s)", s16.Length+uint16(size)-uint16(length))
		return
	}
	used = size + int(s16.Length)
	s16.Value = string(data[size:used])
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (s16 *String16) UnmarshalLittleEndian(data []byte) (used int, err error) {
	var size int = binary.Size(s16.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	s16.Length = binary.LittleEndian.Uint16(data[:size])
	if uint16(length-size) < s16.Length {
		err = fmt.Errorf("Need more %d byte(s)", s16.Length+uint16(size)-uint16(length))
		return
	}
	used = size + int(s16.Length)
	s16.Value = string(data[size:used])
	return
}

// String32 defines a common string type, which max length is math.MaxUint32.
//
// +--....--+-------+--------+   +--------+
// | length |char-0 | char-1 |...| char-n |
// +--....--+-------+--------+   +--------+
type String32 struct {
	Length uint32
	Value  string
}

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (s32 String32) MarshalBigEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(s32.Length))
	binary.BigEndian.PutUint32(mem, s32.Length)
	return append(mem, []byte(s32.Value)...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (s32 String32) MarshalLittleEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(s32.Length))
	binary.LittleEndian.PutUint32(mem, s32.Length)
	return append(mem, []byte(s32.Value)...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (s32 *String32) UnmarshalBigEndian(data []byte) (used int, err error) {
	var size int = binary.Size(s32.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	s32.Length = binary.BigEndian.Uint32(data[:size])
	if uint32(length-size) < s32.Length {
		err = fmt.Errorf("Need more %d byte(s)", s32.Length+uint32(size)-uint32(length))
		return
	}
	used = size + int(s32.Length)
	s32.Value = string(data[size:used])
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (s32 *String32) UnmarshalLittleEndian(data []byte) (used int, err error) {
	var size int = binary.Size(s32.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	s32.Length = binary.LittleEndian.Uint32(data[:size])
	if uint32(length-size) < s32.Length {
		err = fmt.Errorf("Need more %d byte(s)", s32.Length+uint32(size)-uint32(length))
		return
	}
	used = size + int(s32.Length)
	s32.Value = string(data[size:used])
	return
}

// String64 defines a common string type, which max length is math.MaxUint64.
//
// +--....--+-------+--------+   +--------+
// | length |char-0 | char-1 |...| char-n |
// +--....--+-------+--------+   +--------+
type String64 struct {
	Length uint64
	Value  string
}

// MarshalBigEndian implements the BigEndianMarshaler interface.
func (s64 String64) MarshalBigEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(s64.Length))
	binary.BigEndian.PutUint64(mem, s64.Length)
	return append(mem, []byte(s64.Value)...), nil
}

// MarshalLittleEndian implements the LittleEndianMarshaler interface.
func (s64 String64) MarshalLittleEndian() ([]byte, error) {
	var mem []byte = make([]byte, binary.Size(s64.Length))
	binary.LittleEndian.PutUint64(mem, s64.Length)
	return append(mem, []byte(s64.Value)...), nil
}

// UnmarshalBigEndian implements the BigEndianUnmarshaler interface.
func (s64 *String64) UnmarshalBigEndian(data []byte) (used int, err error) {
	var size int = binary.Size(s64.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	s64.Length = binary.BigEndian.Uint64(data[:size])
	if uint64(length-size) < s64.Length {
		err = fmt.Errorf("Need more %d byte(s)", s64.Length+uint64(size)-uint64(length))
		return
	}
	used = size + int(s64.Length)
	s64.Value = string(data[size:used])
	return
}

// UnmarshalLittleEndian implements the LittleEndianUnmarshaler interface.
func (s64 *String64) UnmarshalLittleEndian(data []byte) (used int, err error) {
	var size int = binary.Size(s64.Length)
	var length int = len(data)
	if length < size {
		err = fmt.Errorf("Need more %d byte(s)", size-length)
		return
	}
	s64.Length = binary.LittleEndian.Uint64(data[:size])
	if uint64(length-size) < s64.Length {
		err = fmt.Errorf("Need more %d byte(s)", s64.Length+uint64(size)-uint64(length))
		return
	}
	used = size + int(s64.Length)
	s64.Value = string(data[size:used])
	return
}
