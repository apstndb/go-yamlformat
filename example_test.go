package yamlformat_test

import (
	"fmt"
	"os"

	"github.com/apstndb/go-yamlformat"
)

func ExampleFormat() {
	// Parse format from string
	format, err := yamlformat.ParseFormat("yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(format)
	
	// Check if format is valid
	if yamlformat.FormatJSON.IsValid() {
		fmt.Println("JSON is a valid format")
	}
	
	// Output:
	// yaml
	// JSON is a valid format
}

func ExampleNewEncoderForFormat() {
	data := map[string]interface{}{
		"name": "example",
		"items": []string{"foo", "bar"},
		"count": 42,
	}
	
	// Create YAML encoder
	yamlEncoder := yamlformat.NewEncoderForFormat(os.Stdout, yamlformat.FormatYAML)
	fmt.Println("YAML output:")
	yamlEncoder.Encode(data)
	
	fmt.Println("\nJSON output:")
	// Create JSON encoder
	jsonEncoder := yamlformat.NewEncoderForFormat(os.Stdout, yamlformat.FormatJSON)
	jsonEncoder.Encode(data)
	
	// Output:
	// YAML output:
	// count: 42
	// items:
	// - foo
	// - bar
	// name: example
	//
	// JSON output:
	// {"count": 42, "items": ["foo", "bar"], "name": "example"}
}

func ExampleMarshal() {
	data := map[string]interface{}{
		"message": "Hello, World!",
		"pi":      3.14,
		"whole":   100.0, // Will be converted to 100
	}
	
	yamlBytes, err := yamlformat.Marshal(data)
	if err != nil {
		panic(err)
	}
	
	fmt.Print(string(yamlBytes))
	
	// Output:
	// message: Hello, World!
	// pi: 3.14
	// whole: 100
}

func ExampleMarshalJSON() {
	data := map[string]interface{}{
		"message": "Hello, World!",
		"pi":      3.14,
		"whole":   100.0, // Will be converted to 100
	}
	
	jsonBytes, err := yamlformat.MarshalJSON(data)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(jsonBytes))
	
	// Output:
	// {"message": "Hello, World!", "pi": 3.14, "whole": 100}
}

func ExampleUnmarshal() {
	// YAML input
	yamlData := []byte(`
name: John Doe
age: 30
active: true
`)
	
	var result map[string]interface{}
	err := yamlformat.Unmarshal(yamlData, &result)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Name: %s\n", result["name"])
	fmt.Printf("Age: %v\n", result["age"])
	fmt.Printf("Active: %v\n", result["active"])
	
	// Output:
	// Name: John Doe
	// Age: 30
	// Active: true
}

func ExampleUnmarshal_json() {
	// JSON input also works
	jsonData := []byte(`{"name":"Jane Doe","age":25,"active":false}`)
	
	var result map[string]interface{}
	err := yamlformat.Unmarshal(jsonData, &result)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Name: %s\n", result["name"])
	fmt.Printf("Age: %v\n", result["age"])
	fmt.Printf("Active: %v\n", result["active"])
	
	// Output:
	// Name: Jane Doe
	// Age: 25
	// Active: false
}

func ExampleMarshal_multilineString() {
	data := map[string]interface{}{
		"description": "This is a\nmulti-line\nstring",
	}
	
	yamlBytes, err := yamlformat.Marshal(data)
	if err != nil {
		panic(err)
	}
	
	fmt.Print(string(yamlBytes))
	
	// Output:
	// description: |-
	//   This is a
	//   multi-line
	//   string
}

func ExampleMarshal_customOptions() {
	data := map[string]interface{}{
		"nested": map[string]interface{}{
			"key": "value",
		},
	}
	
	// Marshal with custom indentation (4 spaces)
	// Note: You need to import github.com/goccy/go-yaml for custom options
	// yamlBytes, err := yamlformat.Marshal(data, yaml.Indent(4))
	
	// For this example, we'll just show the default behavior
	yamlBytes, err := yamlformat.Marshal(data)
	if err != nil {
		panic(err)
	}
	
	fmt.Print(string(yamlBytes))
	
	// Output:
	// nested:
	//   key: value
}