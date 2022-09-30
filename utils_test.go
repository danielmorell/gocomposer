package gocomposer

import "testing"

func Test_isArray(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{
			name: "valid array",
			data: []byte("[1, 2]"),
			want: true,
		},
		{
			name: "empty array",
			data: []byte("[]"),
			want: true,
		},
		{
			name: "array in string",
			data: []byte("\"[1, 2]\""),
			want: false,
		},
		{
			name: "number",
			data: []byte("12"),
			want: false,
		},
		{
			name: "zero",
			data: []byte("0"),
			want: false,
		},
		{
			name: "null",
			data: []byte("null"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArray(tt.data); got != tt.want {
				t.Errorf("isArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isObject(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{
			name: "valid object",
			data: []byte("{\"foo\": 1}"),
			want: true,
		},
		{
			name: "empty object",
			data: []byte("{}"),
			want: true,
		},
		{
			name: "object in string",
			data: []byte("\"{\"foo\": 1}\""),
			want: false,
		},
		{
			name: "number",
			data: []byte("12"),
			want: false,
		},
		{
			name: "zero",
			data: []byte("0"),
			want: false,
		},
		{
			name: "null",
			data: []byte("null"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isObject(tt.data); got != tt.want {
				t.Errorf("isObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isString(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{
			name: "valid string",
			data: []byte("\"foo\""),
			want: true,
		},
		{
			name: "empty object",
			data: []byte("\"\""),
			want: true,
		},
		{
			name: "number",
			data: []byte("12"),
			want: false,
		},
		{
			name: "zero",
			data: []byte("0"),
			want: false,
		},
		{
			name: "null",
			data: []byte("null"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isString(tt.data); got != tt.want {
				t.Errorf("isString() = %v, want %v", got, tt.want)
			}
		})
	}
}
