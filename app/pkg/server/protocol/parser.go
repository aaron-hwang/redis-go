package protocol

// This package contains logic for parsing and manipulating RESP protocol data.

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type RespParser struct {
	reader *bufio.Reader
}

func NewRespParser(reader io.Reader) *RespParser {
	return &RespParser{reader: bufio.NewReader(reader)}
}

// Main function.
func (rp *RespParser) Read() (RespVal, error) {
	typ, err := rp.reader.ReadByte()
	if err != nil {
		return RespVal{}, err
	}

	switch typ {
	case SimpleString:
		return rp.ReadSimpleString()
	case BulkString:
		return rp.ReadBulkString()
	case Integer:
		return rp.ReadInteger()
	case Array:
		return rp.ReadArray()
	default:
		return RespVal{}, nil
	}
}

func (rp *RespParser) ReadLine() (line []byte, n int, err error) {
	for {
		b, err := rp.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n++
		line = append(line, b)
		if len(line) >= 2 && line[n-2] == '\r' && line[n-1] == '\n' {
			break
		}
	}

	return line[:len(line)-2], n, nil
}

// To make life easier for myself, type byte should already be parsed by the time these helpers are called.
// For the purpose of handling nested types, all these functions should also parse past the \r\n delimiters.

func (rp *RespParser) ReadSimpleString() (RespVal, error) {
	text, _, err := rp.ReadLine()
	if err != nil {
		return RespVal{}, err
	}

	return RespVal{str: string(text), typ: SimpleString}, nil
}

func (rp *RespParser) ReadSimpleError() (RespVal, error) {
	v, err := rp.ReadSimpleString()
	if err != nil {
		return RespVal{}, err
	}
	v.typ = SimpleError
	return v, nil
}

func (rp *RespParser) ReadBulkString() (RespVal, error) {
	_, err := rp.parseInt()
	if err != nil {
		return RespVal{}, err
	}
	data, _, err := rp.ReadLine()
	if err != nil {
		return RespVal{}, err
	}
	return RespVal{typ: BulkString, str: string(data)}, nil
}

// Helper function: Parses an int literal from a resp message, and returns a go int.
func (rp *RespParser) parseInt() (int64, error) {
	text, _, err := rp.ReadLine()
	if err != nil {
		return 0, err
	}
	// RESP uses base 10 64 bit integers
	val, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (rp *RespParser) ReadInteger() (RespVal, error) {
	sign, err := rp.reader.ReadByte()
	if err != nil {
		return RespVal{}, err
	}
	val, err := rp.parseInt()
	if err != nil {
		return RespVal{}, err
	}

	if sign == '-' {
		val *= -1
	} else {
		if sign != '+' {
			return RespVal{}, errors.New("error when parsing integer, invalid sign byte")
		}
	}
	return RespVal{typ: Integer, num: val}, nil
}

func (rp *RespParser) ReadArray() (RespVal, error) {
	numElements, err := rp.parseInt()
	if err != nil {
		return RespVal{}, err
	}

	result := RespVal{typ: Array, arr: make([]RespVal, 0, numElements)}

	/* Recursively call read on each element of the array, since
	each element could be any resp value type (another array!)
	*/
	for i := int64(0); i < numElements; i++ {
		rval, err := rp.Read()
		// If we encounter any error at all, throw out the whole result. No half baked messes.
		if err != nil {
			return RespVal{}, err
		}
		result.arr = append(result.arr, rval)
	}
	return result, nil
}

func (rp *RespParser) ReadNull() (RespVal, error) {
	_, _, err := rp.ReadLine()
	if err != nil {
		return RespVal{}, err
	}
	return RespVal{typ: Null}, nil
}
