package protocol

// This file has all the logic for representing the RESP protocol.

type RespType = byte

/*
	https://redis.io/docs/latest/develop/reference/protocol-spec/#resp-protocol-description

N.B: RESP used to support representing null types with special cases of bulk or array types.
This is RESP3, so we will not be doing that here.
*/
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
	typ RespType
	num int64
	str string
	arr []RespVal
}
