# go-yamlformat

A Go package for consistent YAML/JSON formatting using goccy/go-yaml.

## Features

- Consistent number handling (preserves int vs float types)
- Unified API for both YAML and JSON encoding/decoding
- Automatic conversion of whole floats to integers (100.0 → 100)
- Proper handling of large integers without scientific notation
- Multi-line strings use literal style (|) by default
- Reusable encoding/decoding options

## Installation

```bash
go get github.com/apstndb/go-yamlformat
```

## Usage

```go
package main

import (
    "fmt"
    "os"
    "github.com/apstndb/go-yamlformat"
)

func main() {
    data := map[string]interface{}{
        "name": "example",
        "count": 42,
        "ratio": 3.14,
        "whole": 100.0,  // Will be encoded as 100
    }
    
    // Marshal to YAML
    yamlBytes, err := yamlformat.Marshal(data)
    if err != nil {
        panic(err)
    }
    fmt.Printf("YAML:\n%s\n", yamlBytes)
    
    // Marshal to JSON
    jsonBytes, err := yamlformat.MarshalJSON(data)
    if err != nil {
        panic(err)
    }
    fmt.Printf("JSON: %s\n", jsonBytes)
    
    // Marshal with custom options
    customYAML, err := yamlformat.Marshal(data, yaml.Indent(4))
    if err != nil {
        panic(err)
    }
    
    // Create encoder for streaming
    encoder := yamlformat.NewEncoderForFormat(os.Stdout, yamlformat.FormatYAML)
    encoder.Encode(data)
}
```

## API

### Types

- `Format`: Represents output format (YAML or JSON)
  - `FormatYAML`: YAML format
  - `FormatJSON`: JSON format

### Functions

- `Marshal(v interface{}, opts ...yaml.EncodeOption) ([]byte, error)`: Marshal to YAML bytes
- `MarshalJSON(v interface{}, opts ...yaml.EncodeOption) ([]byte, error)`: Marshal to JSON bytes
- `Unmarshal(data []byte, v interface{}, opts ...yaml.DecodeOption) error`: Unmarshal from YAML/JSON bytes
- `NewEncoder(w io.Writer, opts ...yaml.EncodeOption) *yaml.Encoder`: Create YAML encoder
- `NewJSONEncoder(w io.Writer, opts ...yaml.EncodeOption) *yaml.Encoder`: Create JSON encoder
- `NewEncoderForFormat(w io.Writer, format Format) *yaml.Encoder`: Create encoder for specified format
- `ParseFormat(s string) (Format, error)`: Parse format string ("yaml" or "json")

### Default Options

#### Encoding (Marshal) Options
- `yaml.UseJSONMarshaler()`: Use JSON marshaling rules for consistency
- `yaml.AutoInt()`: Convert whole floats to integers (100.0 → 100)
- `yaml.UseLiteralStyleIfMultiline(true)`: Use literal style (|) for multi-line strings. Pass `false` to use quoted style instead.

#### Decoding (Unmarshal) Options
- `yaml.UseJSONUnmarshaler()`: Use JSON unmarshaling rules for consistency

Note: To completely remove a default option (e.g., disable AutoInt), you need to use goccy/go-yaml directly as this package always includes the defaults.

## License

MIT