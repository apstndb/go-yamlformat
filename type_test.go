package yamlformat

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAutoIntConversion(t *testing.T) {
	// Test that AutoInt converts whole floats to integers
	input := map[string]interface{}{
		"whole_float": 100.0,
		"actual_float": 3.14,
		"integer": 42,
		"negative_whole": -50.0,
		"large_whole": 1000000.0,
	}
	
	yamlBytes, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	// Unmarshal back to verify types
	var result map[string]interface{}
	err = Unmarshal(yamlBytes, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	
	// Check types
	tests := []struct {
		key      string
		wantType string
		wantValue interface{}
	}{
		{"whole_float", "uint64", uint64(100)},
		{"actual_float", "float64", 3.14},
		{"integer", "uint64", uint64(42)},
		{"negative_whole", "int64", int64(-50)},
		{"large_whole", "uint64", uint64(1000000)},
	}
	
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			got, ok := result[tt.key]
			if !ok {
				t.Fatalf("Key %q not found in result", tt.key)
			}
			
			if diff := cmp.Diff(tt.wantValue, got); diff != "" {
				t.Errorf("Value mismatch for %q (-want +got):\n%s", tt.key, diff)
			}
		})
	}
}

func TestMarshalWithExplicitTypes(t *testing.T) {
	// Test marshaling with specific numeric types
	input := map[string]interface{}{
		"int8_val":    int8(127),
		"int16_val":   int16(32767),
		"int32_val":   int32(2147483647),
		"int64_val":   int64(9223372036854775807),
		"uint8_val":   uint8(255),
		"uint16_val":  uint16(65535),
		"uint32_val":  uint32(4294967295),
		"uint64_val":  uint64(18446744073709551615),
		"float32_val": float32(3.14),
		"float64_val": float64(3.141592653589793),
	}
	
	yamlBytes, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	t.Logf("Marshaled YAML:\n%s", yamlBytes)
	
	// Verify the YAML contains expected values
	yamlStr := string(yamlBytes)
	expectedValues := []string{
		"int8_val: 127",
		"int16_val: 32767",
		"int32_val: 2147483647",
		"int64_val: 9223372036854775807",
		"uint8_val: 255",
		"uint16_val: 65535",
		"uint32_val: 4294967295",
		"uint64_val: 18446744073709551615",
	}
	
	for _, expected := range expectedValues {
		if !contains(yamlStr, expected) {
			t.Errorf("Expected %q in YAML output", expected)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		   len(s) >= len(substr) && contains(s[1:], substr)
}