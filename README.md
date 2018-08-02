# bin #
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

## Install ##
```
go get -u github.com/whiler/bin
```

## Examples ##
Let's send one SOCKS 5 greeting packt to remote.
```
type Reply struct {
	Ver byte
	Method byte
}

reply := Reply{Ver: 5, Method: 2}
err := bin.MarshalBigEndianTo(remote, reply)
```

## Supported types ##
`fixed-size types` including `bool`, `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `complex64`, `complex128` and an array or struct containing only fixed-size types.

|                           | fixed-size types | string | marshaler               | unmarshaler             |
|---------------------------|------------------|--------|-------------------------|-------------------------|
| MarshalBigEndian          | yes              | yes    | BigEndianMarshaler      |                         |
| MarshalBigEndianTo        | yes              | yes    | BigEndianMarshaler      |                         |
| MarshalLittleEndian       | yes              | yes    | LittleEndianMarshaler   |                         |
| MarshalLittleEndianTo     | yes              | yes    | LittleEndianMarshaler   |                         |
| UnmarshalBigEndian        | yes              | no     |                         | BigEndianUnmarshaler    |
| UnmarshalBigEndianFrom    | yes              | no     |                         | BigEndianUnmarshaler    |
| UnmarshalLittleEndian     | yes              | no     |                         | LittleEndianUnmarshaler |
| UnmarshalLittleEndianFrom | yes              | no     |                         | LittleEndianUnmarshaler |

