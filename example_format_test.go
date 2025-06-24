package yamlformat_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/apstndb/go-yamlformat"
)

func ExampleParseFormat() {
	// Parse format from user input
	formats := []string{"yaml", "JSON", "YAML", "json", "invalid"}
	
	for _, input := range formats {
		format, err := yamlformat.ParseFormat(input)
		if err != nil {
			fmt.Printf("%s: error - %v\n", input, err)
		} else {
			fmt.Printf("%s: %v\n", input, format)
		}
	}
	
	// Output:
	// yaml: yaml
	// JSON: json
	// YAML: yaml
	// json: json
	// invalid: error - invalid format: invalid (valid: yaml, json)
}

func ExampleFormat_IsValid() {
	// Check if a format is valid
	fmt.Println("FormatYAML valid:", yamlformat.FormatYAML.IsValid())
	fmt.Println("FormatJSON valid:", yamlformat.FormatJSON.IsValid())
	fmt.Println("Invalid format valid:", yamlformat.Format("xml").IsValid())
	
	// Output:
	// FormatYAML valid: true
	// FormatJSON valid: true
	// Invalid format valid: false
}

// ExampleNewEncoderForFormat_dynamicFormat shows how to dynamically select format
func ExampleNewEncoderForFormat_dynamicFormat() {
	data := map[string]interface{}{
		"status": "success",
		"code":   200,
	}
	
	// Simulate command-line flag or config
	outputFormats := []yamlformat.Format{yamlformat.FormatYAML, yamlformat.FormatJSON}
	
	for _, format := range outputFormats {
		var buf bytes.Buffer
		encoder := yamlformat.NewEncoderForFormat(&buf, format)
		
		if err := encoder.Encode(data); err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("Format %s:\n%s\n", format, buf.String())
	}
	
	// Output:
	// Format yaml:
	// code: 200
	// status: success
	//
	// Format json:
	// {"code": 200, "status": "success"}
}