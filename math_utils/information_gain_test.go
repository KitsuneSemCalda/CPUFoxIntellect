package mathutils

import (
	"testing"
)

func TestCalcInformationGain(t *testing.T) {
	tests := []struct {
		name           string
		data           []DataPoint
		attributeIndex int
		expected       float32
	}{
		{
			name: "Test 1: All elements are of the same class",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1, 0.2}},
				{Class: 1, Attributes: []float32{0.3, 0.4}},
				{Class: 1, Attributes: []float32{0.5, 0.6}},
				{Class: 1, Attributes: []float32{0.7, 0.8}},
			},
			attributeIndex: 0,
			expected:       0.0,
		},
		{
			name: "Test 2: All elements are of the same class, different attributes",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1, 0.2}},
				{Class: 1, Attributes: []float32{0.3, 0.4}},
				{Class: 1, Attributes: []float32{0.5, 0.6}},
				{Class: 1, Attributes: []float32{0.7, 0.8}},
			},
			attributeIndex: 1,
			expected:       0.0,
		},
		{
			name: "Test 3: Elements are of different classes, same attributes",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1, 0.2}},
				{Class: 2, Attributes: []float32{0.1, 0.2}},
				{Class: 3, Attributes: []float32{0.1, 0.2}},
				{Class: 4, Attributes: []float32{0.1, 0.2}},
			},
			attributeIndex: 0,
			expected:       0.0,
		},
		{
			name: "Test 4: Some elements are of the same class, different attributes",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1, 0.2}},
				{Class: 2, Attributes: []float32{0.3, 0.4}},
				{Class: 2, Attributes: []float32{0.5, 0.6}},
				{Class: 3, Attributes: []float32{0.7, 0.8}},
			},
			attributeIndex: 1,
			expected:       1.5,
		},
		{
			name: "Test 5: All elements are of the same class, one attribute",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1}},
				{Class: 1, Attributes: []float32{0.2}},
				{Class: 1, Attributes: []float32{0.3}},
				{Class: 1, Attributes: []float32{0.4}},
			},
			attributeIndex: 0,
			expected:       0.0,
		},
		{
			name: "Test 6: Elements are of different classes, one attribute",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1}},
				{Class: 2, Attributes: []float32{0.2}},
				{Class: 3, Attributes: []float32{0.3}},
				{Class: 4, Attributes: []float32{0.4}},
			},
			attributeIndex: 0,
			expected:       2.0,
		},
		{
			name: "Test 7: Some elements are of the same class, one attribute",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{0.1}},
				{Class: 2, Attributes: []float32{0.2}},
				{Class: 2, Attributes: []float32{0.3}},
				{Class: 3, Attributes: []float32{0.4}},
			},
			attributeIndex: 0,
			expected:       1.5,
		},
		{
			name: "Test 8: All elements are of the same class, no attributes",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{}},
				{Class: 1, Attributes: []float32{}},
				{Class: 1, Attributes: []float32{}},
				{Class: 1, Attributes: []float32{}},
			},
			attributeIndex: 0,
			expected:       0.0,
		},
		{
			name: "Test 9: Elements are of different classes, no attributes",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{}},
				{Class: 2, Attributes: []float32{}},
				{Class: 3, Attributes: []float32{}},
				{Class: 4, Attributes: []float32{}},
			},
			attributeIndex: 0,
			expected:       0.0,
		},
		{
			name: "Test 10: Some elements are of the same class, no attributes",
			data: []DataPoint{
				{Class: 1, Attributes: []float32{}},
				{Class: 2, Attributes: []float32{}},
				{Class: 2, Attributes: []float32{}},
				{Class: 3, Attributes: []float32{}},
			},
			attributeIndex: 0,
			expected:       0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalcInformationGain(tt.data, tt.attributeIndex)
			if actual != tt.expected {
				t.Errorf("CalcInformationGain(%v, %v) = %v; want %v", tt.data, tt.attributeIndex, actual, tt.expected)
			}
		})
	}
}
