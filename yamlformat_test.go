package yamlformat

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		opts    []yaml.EncodeOption
		want    string
		wantErr bool
	}{
		{
			name: "simple map",
			input: map[string]interface{}{
				"name":  "test",
				"value": 42,
			},
			want: "name: test\nvalue: 42\n",
		},
		{
			name: "auto int conversion",
			input: map[string]interface{}{
				"whole": 100.0,
				"float": 3.14,
			},
			want: "float: 3.14\nwhole: 100\n",
		},
		{
			name: "multi-line string with literal style",
			input: map[string]interface{}{
				"description": "This is a\nmulti-line\nstring",
			},
			want: "description: |-\n  This is a\n  multi-line\n  string\n",
		},
		{
			name: "multi-line string with quoted style",
			input: map[string]interface{}{
				"description": "This is a\nmulti-line\nstring",
			},
			opts: []yaml.EncodeOption{yaml.UseLiteralStyleIfMultiline(false)},
			want: "description: |-\n  This is a\n  multi-line\n  string\n",
		},
		{
			name: "custom indent",
			input: map[string]interface{}{
				"nested": map[string]interface{}{
					"key": "value",
				},
			},
			opts: []yaml.EncodeOption{yaml.Indent(4)},
			want: "nested:\n    key: value\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("Marshal() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		opts    []yaml.EncodeOption
		want    string
		wantErr bool
	}{
		{
			name: "simple map",
			input: map[string]interface{}{
				"name":  "test",
				"value": 42,
			},
			want: `{"name": "test", "value": 42}` + "\n",
		},
		{
			name: "auto int conversion",
			input: map[string]interface{}{
				"whole": 100.0,
				"float": 3.14,
			},
			want: `{"float": 3.14, "whole": 100}` + "\n",
		},
		{
			name: "nested structure",
			input: map[string]interface{}{
				"user": map[string]interface{}{
					"id":   123,
					"name": "John",
				},
			},
			want: `{"user": {"id": 123, "name": "John"}}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalJSON(tt.input, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("MarshalJSON() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    map[string]interface{}
		opts    []yaml.DecodeOption
		wantErr bool
	}{
		{
			name:  "YAML input with integer",
			input: "name: test\nvalue: 42\n",
			want: map[string]interface{}{
				"name":  "test",
				"value": uint64(42),
			},
		},
		{
			name:  "JSON input with integer",
			input: `{"name":"test","value":42}`,
			want: map[string]interface{}{
				"name":  "test",
				"value": uint64(42),
			},
		},
		{
			name:  "YAML with float",
			input: "pi: 3.14\nwhole: 100.0\n",
			want: map[string]interface{}{
				"pi":    3.14,
				"whole": float64(100.0),
			},
		},
		{
			name:  "multi-line YAML",
			input: "description: |\n  This is a\n  multi-line\n  string\n",
			want: map[string]interface{}{
				"description": "This is a\nmulti-line\nstring\n",
			},
		},
		{
			name:    "invalid input",
			input:   "[1, 2, }",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got map[string]interface{}
			err := Unmarshal([]byte(tt.input), &got, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("Unmarshal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNewEncoder(t *testing.T) {
	var buf bytes.Buffer
	encoder := NewEncoder(&buf)
	
	data := map[string]interface{}{
		"test": "value",
	}
	
	err := encoder.Encode(data)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	
	want := "test: value\n"
	if buf.String() != want {
		t.Errorf("Encode() = %q, want %q", buf.String(), want)
	}
}

func TestNewJSONEncoder(t *testing.T) {
	var buf bytes.Buffer
	encoder := NewJSONEncoder(&buf)
	
	data := map[string]interface{}{
		"test": "value",
	}
	
	err := encoder.Encode(data)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	
	want := `{"test": "value"}`
	if strings.TrimSpace(buf.String()) != want {
		t.Errorf("Encode() = %q, want %q", strings.TrimSpace(buf.String()), want)
	}
}

func TestFormatMarshal(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		input  map[string]interface{}
		want   string
	}{
		{
			name:   "YAML format",
			format: FormatYAML,
			input:  map[string]interface{}{"key": "value"},
			want:   "key: value\n",
		},
		{
			name:   "JSON format",
			format: FormatJSON,
			input:  map[string]interface{}{"key": "value"},
			want:   `{"key": "value"}` + "\n",
		},
		{
			name:   "invalid format defaults to YAML",
			format: Format("invalid"),
			input:  map[string]interface{}{"key": "value"},
			want:   "key: value\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.format.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}
			
			if string(got) != tt.want {
				t.Errorf("Marshal() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestFormatNewEncoder(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		input  map[string]interface{}
		want   string
	}{
		{
			name:   "YAML format",
			format: FormatYAML,
			input:  map[string]interface{}{"key": "value"},
			want:   "key: value\n",
		},
		{
			name:   "JSON format",
			format: FormatJSON,
			input:  map[string]interface{}{"key": "value"},
			want:   `{"key": "value"}`,
		},
		{
			name:   "invalid format defaults to YAML",
			format: Format("invalid"),
			input:  map[string]interface{}{"key": "value"},
			want:   "key: value\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := tt.format.NewEncoder(&buf)
			
			err := encoder.Encode(tt.input)
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}
			
			got := strings.TrimSpace(buf.String())
			want := strings.TrimSpace(tt.want)
			if got != want {
				t.Errorf("Encode() = %q, want %q", got, want)
			}
		})
	}
}

func TestNewEncoderForFormat(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		input  map[string]interface{}
		want   string
	}{
		{
			name:   "YAML format",
			format: FormatYAML,
			input:  map[string]interface{}{"key": "value"},
			want:   "key: value\n",
		},
		{
			name:   "JSON format",
			format: FormatJSON,
			input:  map[string]interface{}{"key": "value"},
			want:   `{"key": "value"}`,
		},
		{
			name:   "invalid format defaults to YAML",
			format: Format("invalid"),
			input:  map[string]interface{}{"key": "value"},
			want:   "key: value\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := NewEncoderForFormat(&buf, tt.format)
			
			err := encoder.Encode(tt.input)
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}
			
			got := strings.TrimSpace(buf.String())
			want := strings.TrimSpace(tt.want)
			if got != want {
				t.Errorf("Encode() = %q, want %q", got, want)
			}
		})
	}
}

func TestParseFormat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Format
		wantErr bool
	}{
		{
			name:  "yaml lowercase",
			input: "yaml",
			want:  FormatYAML,
		},
		{
			name:  "json lowercase",
			input: "json",
			want:  FormatJSON,
		},
		{
			name:  "YAML uppercase",
			input: "YAML",
			want:  FormatYAML,
		},
		{
			name:  "JSON uppercase",
			input: "JSON",
			want:  FormatJSON,
		},
		{
			name:    "invalid format",
			input:   "xml",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFormat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}


