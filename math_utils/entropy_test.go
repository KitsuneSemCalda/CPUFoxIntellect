package mathutils

import (
	"testing"
)

func TestCalcEntropy(t *testing.T) {
	tests := []struct {
		name     string
		data     []DataPoint
		expected float32
	}{
		{
			name:     "Test 1: All elements are of the same class",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1, 0.2}}, {Class: 1, Attributes: []float32{0.3, 0.4}}, {Class: 1, Attributes: []float32{0.5, 0.6}}, {Class: 1, Attributes: []float32{0.7, 0.8}}},
			expected: 0.0,
		},
		{
			name:     "Test 2: Elements are of different classes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1, 0.2}}, {Class: 2, Attributes: []float32{0.3, 0.4}}, {Class: 3, Attributes: []float32{0.5, 0.6}}, {Class: 4, Attributes: []float32{0.7, 0.8}}},
			expected: 2.0,
		},
		{
			name:     "Test 3: Some elements are of the same class",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1, 0.2}}, {Class: 2, Attributes: []float32{0.3, 0.4}}, {Class: 2, Attributes: []float32{0.5, 0.6}}, {Class: 3, Attributes: []float32{0.7, 0.8}}},
			expected: 1.5,
		},
		{
			name:     "Test 4: All elements are of the same class, different attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1, 0.2}}, {Class: 1, Attributes: []float32{0.3, 0.4}}, {Class: 1, Attributes: []float32{0.5, 0.6}}, {Class: 1, Attributes: []float32{0.7, 0.8}}},
			expected: 0.0,
		},
		{
			name:     "Test 5: Elements are of different classes, same attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1, 0.2}}, {Class: 2, Attributes: []float32{0.1, 0.2}}, {Class: 3, Attributes: []float32{0.1, 0.2}}, {Class: 4, Attributes: []float32{0.1, 0.2}}},
			expected: 2.0,
		},
		{
			name:     "Test 6: Some elements are of the same class, different attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1, 0.2}}, {Class: 2, Attributes: []float32{0.3, 0.4}}, {Class: 2, Attributes: []float32{0.5, 0.6}}, {Class: 3, Attributes: []float32{0.7, 0.8}}},
			expected: 1.5,
		},
		{
			name:     "Test 7: All elements are of the same class, one attribute",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1}}, {Class: 1, Attributes: []float32{0.2}}, {Class: 1, Attributes: []float32{0.3}}, {Class: 1, Attributes: []float32{0.4}}},
			expected: 0.0,
		},
		{
			name:     "Test 8: Elements are of different classes, one attribute",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1}}, {Class: 2, Attributes: []float32{0.2}}, {Class: 3, Attributes: []float32{0.3}}, {Class: 4, Attributes: []float32{0.4}}},
			expected: 2.0,
		},
		{
			name:     "Test 9: Some elements are of the same class, one attribute",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1}}, {Class: 2, Attributes: []float32{0.2}}, {Class: 2, Attributes: []float32{0.3}}, {Class: 3, Attributes: []float32{0.4}}},
			expected: 1.5,
		},
		{
			name:     "Test 10: All elements are of the same class, no attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{}}, {Class: 1, Attributes: []float32{}}, {Class: 1, Attributes: []float32{}}, {Class: 1, Attributes: []float32{}}},
			expected: 0.0,
		},
		{
			name:     "Test 11: Elements are of different classes, no attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{}}, {Class: 2, Attributes: []float32{}}, {Class: 3, Attributes: []float32{}}, {Class: 4, Attributes: []float32{}}},
			expected: 2.0,
		},
		{
			name:     "Test 12: Some elements are of the same class, no attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{}}, {Class: 2, Attributes: []float32{}}, {Class: 2, Attributes: []float32{}}, {Class: 3, Attributes: []float32{}}},
			expected: 1.5,
		},
		{
			name:     "Test 13: All elements are of the same class, different number of attributes",
			data:     []DataPoint{{Class: 1, Attributes: []float32{0.1}}, {Class: 1, Attributes: []float32{0.2, 0.3}}, {Class: 1, Attributes: []float32{0.4, 0.5, 0.6}}, {Class: 1, Attributes: []float32{0.7, 0.8, 0.9, 1.0}}},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalcEntropy(tt.data)
			if actual != tt.expected {
				t.Errorf("CalcEntropy(%v) = %v; want %v", tt.data, actual, tt.expected)
			}
		})
	}
}
