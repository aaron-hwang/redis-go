package protocol

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestRespParser(t *testing.T) {
	parser := NewRespParser(strings.NewReader("*1\r\n$4\r\nping\r\n"))
	val, err := parser.Read()
	if err != nil {
		t.Fatalf("Parser came across an error where none expected %s", err)
	}
	fmt.Printf("t: %v\n", val)
}

func TestRandomTest(t *testing.T) {
	parser := NewRespParser(bufio.NewReader(strings.NewReader("*1\r\n$4\r\nping\r\n")))
	val, _, err := parser.ReadLine()
	if err != nil {
		t.Fatalf("Parser came across an error where none expected %s", err)
	}
	t.Logf("test: %v", string(val))
}
