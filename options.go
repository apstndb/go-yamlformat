package yamlformat

import "github.com/goccy/go-yaml"

// DefaultMarshalOptions returns a copy of the default marshal options
func DefaultMarshalOptions() []yaml.EncodeOption {
	return append([]yaml.EncodeOption{}, MarshalOptions...)
}

// DefaultUnmarshalOptions returns a copy of the default unmarshal options
func DefaultUnmarshalOptions() []yaml.DecodeOption {
	return append([]yaml.DecodeOption{}, UnmarshalOptions...)
}

// WithMarshalOptions creates a new set of options by appending to defaults
func WithMarshalOptions(opts ...yaml.EncodeOption) []yaml.EncodeOption {
	return append(DefaultMarshalOptions(), opts...)
}

// WithUnmarshalOptions creates a new set of options by appending to defaults
func WithUnmarshalOptions(opts ...yaml.DecodeOption) []yaml.DecodeOption {
	return append(DefaultUnmarshalOptions(), opts...)
}

// WithJSONMarshalOptions creates a new set of JSON encoding options
func WithJSONMarshalOptions(opts ...yaml.EncodeOption) []yaml.EncodeOption {
	allOpts := DefaultMarshalOptions()
	allOpts = append(allOpts, yaml.JSON())
	allOpts = append(allOpts, opts...)
	return allOpts
}