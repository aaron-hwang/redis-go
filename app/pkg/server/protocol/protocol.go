package protocol

// This file has all the logic for parsing and representing the RESP protocol.

type RespType = byte

const (
	SimpleString RespType = '+'
	SimpleError  RespType = '-'
	Integer      RespType = ':'
	BulkString   RespType = '$'
	// Special value not used in normal resp serialization, placeholder to internally tell this is a null bulk string
	NullBulkString RespType = 'N'
	Array          RespType = '*'
	Null           RespType = '_'
	Boolean        RespType = '#'
	Double         RespType = ','
	BigNum         RespType = '('
	BulkError      RespType = '!'
	VerbatimString RespType = '='
	Map            RespType = '%'
	Attribute      RespType = '.'
	Set            RespType = '~'
	Push           RespType = '>'
)

// Struct for representing an in-memory RESP protocol value.
type RespVal struct {
	typ_ RespType
	int_ int
	str_ string
	arr  []RespVal
}
