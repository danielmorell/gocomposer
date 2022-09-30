package gocomposer

import (
	"encoding/json"
	is2 "github.com/matryer/is"
	"testing"
)

func TestStringOrBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		output string
		value  StringOrBool
	}{
		{
			name:   `BoolFalse`,
			output: `false`,
			value: StringOrBool{
				isBool:      true,
				boolValue:   false,
				stringValue: "",
			},
		},
		{
			name:   `BoolTrue`,
			output: `true`,
			value: StringOrBool{
				isBool:      true,
				boolValue:   true,
				stringValue: "",
			},
		},
		{
			name:   `StringSimple`,
			output: `"simple"`,
			value: StringOrBool{
				isBool:      false,
				boolValue:   false,
				stringValue: "simple",
			},
		},
		{
			name:   `StringEscape`,
			output: `"Escape this \" if you can."`,
			value: StringOrBool{
				isBool:      false,
				boolValue:   false,
				stringValue: `Escape this " if you can.`,
			},
		},
		{
			name:   `Empty`,
			output: `""`,
			value: StringOrBool{
				isBool:      false,
				boolValue:   false,
				stringValue: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := is2.New(t)

			result, err := json.Marshal(test.value)

			is.NoErr(err)
			is.Equal(string(result), test.output)
		})
	}
}

func TestStringOrBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		value StringOrBool
	}{
		{
			name:  `BoolFalse`,
			input: `false`,
			value: StringOrBool{
				isBool:      true,
				boolValue:   false,
				stringValue: "",
			},
		},
		{
			name:  `BoolTrue`,
			input: `true`,
			value: StringOrBool{
				isBool:      true,
				boolValue:   true,
				stringValue: "",
			},
		},
		{
			name:  `StringSimple`,
			input: `"simple"`,
			value: StringOrBool{
				isBool:      false,
				boolValue:   false,
				stringValue: "simple",
			},
		},
		{
			name:  `StringEscape`,
			input: `"Escape this \" if you can."`,
			value: StringOrBool{
				isBool:      false,
				boolValue:   false,
				stringValue: `Escape this \" if you can.`,
			},
		},
		{
			name:  `Empty`,
			input: `""`,
			value: StringOrBool{
				isBool:      false,
				boolValue:   false,
				stringValue: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := is2.New(t)

			result := StringOrBool{}
			err := json.Unmarshal([]byte(test.input), &result)

			is.NoErr(err)
			is.Equal(result, test.value)
		})
	}
}
