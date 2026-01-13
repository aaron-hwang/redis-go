package protocol

import (
	"strings"
	"testing"
)

func TestRespParser(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectType RespType
		wantStr    string
		wantNum    int64
	}{
		{
			name:       "Simple string",
			input:      "+hello world\r\n",
			expectType: SimpleString,
			wantStr:    "hello world",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := NewRespParser(strings.NewReader(test.input))
			val, err := parser.Read()
			if err != nil {
				t.Fatalf("error in test %v: %v", test.name, err)
			}

			if val.typ != test.expectType {
				t.Fatalf("expected type of %v, got %v", val.typ, test.expectType)
			}

			if test.wantStr != "" && val.str != test.wantStr {
				t.Fatalf("expected %q, got %q", test.wantStr, val.str)
			}

			if test.wantNum != 0 && val.num != test.wantNum {
				t.Fatalf("expected %d, got %d", test.wantNum, val.num)
			}
		})
	}
}
