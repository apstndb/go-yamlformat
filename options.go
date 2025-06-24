package yamlformat

import "github.com/goccy/go-yaml"

// defaultMarshalOptions returns a copy of the default marshal options
func defaultMarshalOptions() []yaml.EncodeOption {
	return append([]yaml.EncodeOption{}, marshalOptions...)
}

// defaultUnmarshalOptions returns a copy of the default unmarshal options
func defaultUnmarshalOptions() []yaml.DecodeOption {
	return append([]yaml.DecodeOption{}, unmarshalOptions...)
}

// WithMarshalOptions creates a new set of options by appending to defaults
func WithMarshalOptions(opts ...yaml.EncodeOption) []yaml.EncodeOption {
	return append(defaultMarshalOptions(), opts...)
}

// WithUnmarshalOptions creates a new set of options by appending to defaults
func WithUnmarshalOptions(opts ...yaml.DecodeOption) []yaml.DecodeOption {
	return append(defaultUnmarshalOptions(), opts...)
}