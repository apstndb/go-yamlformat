// Package yamlformat provides consistent YAML/JSON formatting utilities using goccy/go-yaml.
// It ensures proper number handling and provides unified APIs for both encoding and decoding.
package yamlformat

import (
	"fmt"
	"io"
	"strings"

	"github.com/goccy/go-yaml"
)

// Format represents the output format for structured data
type Format string

const (
	FormatYAML Format = "yaml"
	FormatJSON Format = "json"
)

// Common options for consistent behavior
var (
	// MarshalOptions are the default options for marshaling
	MarshalOptions = []yaml.EncodeOption{
		yaml.UseJSONMarshaler(),
		yaml.AutoInt(),
		yaml.UseLiteralStyleIfMultiline(true),
	}
	
	// UnmarshalOptions are the default options for unmarshaling
	UnmarshalOptions = []yaml.DecodeOption{
		yaml.UseJSONUnmarshaler(),
	}
)

// IsValid checks if the format is supported
func (f Format) IsValid() bool {
	return f == FormatYAML || f == FormatJSON
}

// ParseFormat parses a string into a Format
func ParseFormat(s string) (Format, error) {
	format := Format(strings.ToLower(s))
	if !format.IsValid() {
		return "", fmt.Errorf("invalid format: %s (valid: yaml, json)", s)
	}
	return format, nil
}

// Marshal marshals data to YAML bytes using consistent options
func Marshal(v interface{}, opts ...yaml.EncodeOption) ([]byte, error) {
	allOpts := append([]yaml.EncodeOption{}, MarshalOptions...)
	allOpts = append(allOpts, opts...)
	return yaml.MarshalWithOptions(v, allOpts...)
}

// MarshalJSON marshals data to JSON bytes
func MarshalJSON(v interface{}, opts ...yaml.EncodeOption) ([]byte, error) {
	allOpts := append([]yaml.EncodeOption{}, MarshalOptions...)
	allOpts = append(allOpts, yaml.JSON())
	allOpts = append(allOpts, opts...)
	return yaml.MarshalWithOptions(v, allOpts...)
}

// Unmarshal unmarshals YAML/JSON bytes using consistent options
func Unmarshal(data []byte, v interface{}, opts ...yaml.DecodeOption) error {
	allOpts := append([]yaml.DecodeOption{}, UnmarshalOptions...)
	allOpts = append(allOpts, opts...)
	return yaml.UnmarshalWithOptions(data, v, allOpts...)
}

// NewEncoder creates a new YAML encoder with consistent options
func NewEncoder(w io.Writer, opts ...yaml.EncodeOption) *yaml.Encoder {
	allOpts := append([]yaml.EncodeOption{}, MarshalOptions...)
	allOpts = append(allOpts, opts...)
	return yaml.NewEncoder(w, allOpts...)
}

// NewJSONEncoder creates a new JSON encoder with consistent options
func NewJSONEncoder(w io.Writer, opts ...yaml.EncodeOption) *yaml.Encoder {
	allOpts := append([]yaml.EncodeOption{}, MarshalOptions...)
	allOpts = append(allOpts, yaml.JSON())
	allOpts = append(allOpts, opts...)
	return yaml.NewEncoder(w, allOpts...)
}

// NewEncoderForFormat creates a new encoder for the specified format
func NewEncoderForFormat(w io.Writer, format Format) *yaml.Encoder {
	switch format {
	case FormatJSON:
		return NewJSONEncoder(w)
	case FormatYAML:
		return NewEncoder(w)
	default:
		return NewEncoder(w) // Default to YAML
	}
}