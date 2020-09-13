package main

import (
	"bytes"
	"testing"
)

func TestJudge(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string // "Yes" or "No"
	}{
		{
			name:  "example1",
			input: "123456789",
			want:  "Yes\n",
		},
		{
			name:  "example2",
			input: "0",
			want:  "Yes\n",
		},
		{
			name:  "example3",
			input: "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280",
			want:  "No\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			input := bytes.NewBufferString(tt.input)
			judge(input, buf)
			if got := buf.String(); got != tt.want {
				t.Errorf("want:%v got:%v", tt.want, got)
			}
		})
	}
}
