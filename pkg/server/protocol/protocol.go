package protocol

type RespType int

const (
	SimpleString RespType = iota
	Error
	Integer
	BulkString
	Array
)
