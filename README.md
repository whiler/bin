## bin ##
[![Build Status](https://travis-ci.org/whiler/bin.svg?branch=master)](https://travis-ci.org/whiler/bin) [![Coverage Status](https://coveralls.io/repos/github/whiler/bin/badge.svg)](https://coveralls.io/github/whiler/bin) [![Go Report Card](https://goreportcard.com/badge/whiler/bin)](https://goreportcard.com/report/whiler/bin) [![GitHub license](https://img.shields.io/github/license/whiler/bin.svg)](https://github.com/whiler/bin/blob/master/LICENSE)

Package bin implements encoding and decoding of binary data.
The mapping between binary and Go values is described in the documentation for the Marshal and Unmarshal functions.

bin can encode and decode fixed-size values by default.
A fixed-size value is either a fixed-size arithmetic type (bool, int8, uint8, int16, float32, complex64, ...) or an array or struct containing only fixed-size values.

bin can encode and decode variable-length values by implementing the Marshaler and Unmarshaler interface.

bin is designed for encoding and decoding network protocol packets.

The main features are:
- variable-length types support
- omit field support
- byte order support

### Install ###
```
go get -u github.com/whiler/bin
```

### Examples ###
#### variable-length types ####
Parse one SOCKS 5 greeting packt from remote.
```
type ByteSliceType []byte

func (bs *ByteSliceType) UnmarshalBigEndian(data []byte) (used int, err error) {
	size := len(data)
	if size == 0 {
		err = fmt.Errorf("Empty Data")
		return
	}
	length := int(data[0])
	if size < length+1 {
		err = fmt.Errorf("Need more %d byte(s)", length+1-size)
		return
	}
	*bs = make([]byte, length)
	copy(*bs, data[1:length+1])
	used = length + 1
	return
}

type Request struct {
	Ver     byte
	Methods *ByteSliceType
}

req := Request{}
err := bin.UnmarshalBigEndianFrom(remote, &req)
```

#### fixed-size types ####
Send one SOCKS 5 greeting packt to remote.
```
type Reply struct {
	Ver byte
	Method byte
}

reply := Reply{Ver: 5, Method: 2}
err := bin.MarshalBigEndianTo(remote, reply)
```

#### omit and reorder field ####
Omit some fields and reorder the field by tag.
```
type Request struct {
	Ver     byte         `bin:"0"`
	Cmd     byte         `bin:"1"`
	TrackID int          `bin:"-"` // omit this field
	Rsv     byte         `bin:"2"` // reorder this field
	Dst     *AddressType `bin:"3"` // reorder this field
}
```

### Supported types ###
`fixed-size types` including `bool`, `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `complex64`, `complex128` and an array or struct containing only fixed-size types.

The **int** and **uint** types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. They are not `fixed-size types`.

| function                  | fixed-size types | string | marshaler               | unmarshaler             |
|---------------------------|------------------|--------|-------------------------|-------------------------|
| MarshalBigEndian          | yes              | yes    | BigEndianMarshaler      |                         |
| MarshalBigEndianTo        | yes              | yes    | BigEndianMarshaler      |                         |
| MarshalLittleEndian       | yes              | yes    | LittleEndianMarshaler   |                         |
| MarshalLittleEndianTo     | yes              | yes    | LittleEndianMarshaler   |                         |
| UnmarshalBigEndian        | yes              | no     |                         | BigEndianUnmarshaler    |
| UnmarshalBigEndianFrom    | yes              | no     |                         | BigEndianUnmarshaler    |
| UnmarshalLittleEndian     | yes              | no     |                         | LittleEndianUnmarshaler |
| UnmarshalLittleEndianFrom | yes              | no     |                         | LittleEndianUnmarshaler |

#### common types ####
| type    | definition                               |
|---------|------------------------------------------|
| Bytes8  | byte slice type, which max length is 255 |
| String8 | string type, which max length is 255     |

### struct tag ###
tag syntax: `bin:"-"` or `bin:"[0-9]+"`

omit one field while marshaling/unmarshaling with tag `bin:"-"`.

the order of fields in one struct follows the rules below:
- starts at 0
- increases one by one
- no hole
- no repetition
