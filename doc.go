/*
Package bin implements encoding and decoding of binary data.
The mapping between binary and Go values is described in the documentation for the Marshal and Unmarshal functions.

bin can encode and decode fixed-size values by default.
A fixed-size value is either a fixed-size arithmetic type (bool, int8, uint8, int16, float32, complex64, ...) or an array or struct containing only fixed-size values.

bin can encode and decode variable-length values by implementing the Marshaler and Unmarshaler interface.

bin is designed for encoding and decoding network protocol packets.
*/
package bin
