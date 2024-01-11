package mathutils

import (
	"reflect"
	"testing"
)

func TestMethods(t *testing.T) {
	tests := []struct {
		name       string
		class      int
		attributes []float32
		newAttr    float32
		newClass   int
	}{
		{
			name:       "Test 1: Initial class 1, adding attribute 4.0, updating class to 2",
			class:      1,
			attributes: []float32{1.0, 2.0, 3.0},
			newAttr:    4.0,
			newClass:   2,
		},
		{
			name:       "Test 2: Initial class -2, adding attribute -2.0, updating class to -3",
			class:      -2,
			attributes: []float32{-4.0, -5.0, -6.0},
			newAttr:    -2.0,
			newClass:   -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := NewDataPoint(tt.class, tt.attributes)

			expectedClass := tt.class
			if expectedClass < 0 {
				expectedClass *= -1
			}
			if dp.Class != expectedClass {
				t.Errorf("Expected class %d, got %d", expectedClass, dp.Class)
			}

			expectedAttributes := make([]float32, len(tt.attributes))
			copy(expectedAttributes, tt.attributes)
			for i := range expectedAttributes {
				if expectedAttributes[i] < 0 {
					expectedAttributes[i] *= -1
				}
			}
			if !reflect.DeepEqual(dp.Attributes, expectedAttributes) {
				t.Errorf("Expected attributes %v, got %v", expectedAttributes, dp.Attributes)
			}

			expectedNewAttr := tt.newAttr
			if expectedNewAttr < 0 {
				expectedNewAttr *= -1
			}
			dp.AddAttribute(tt.newAttr)
			if !reflect.DeepEqual(dp.Attributes, append(expectedAttributes, expectedNewAttr)) {
				t.Errorf("Expected attributes %v, got %v", append(expectedAttributes, expectedNewAttr), dp.Attributes)
			}

			expectedNewClass := tt.newClass
			if expectedNewClass < 0 {
				expectedNewClass *= -1
			}
			dp.UpdateClass(tt.newClass)
			if dp.Class != expectedNewClass {
				t.Errorf("Expected class %d, got %d", expectedNewClass, dp.Class)
			}
		})
	}
}
