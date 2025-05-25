package constant

import "testing"

func TestGetPoint(t *testing.T) {
	tests := []struct {
		name       string
		value      float64
		thresholds []float64
		expected   int
	}{
		{"thresholds without -1", 10, []float64{5, 10, 15}, 1},
		{"thresholds with -1", 10, []float64{5, -1, 15}, 2},
		{"value less than or equal to first threshold", 5, []float64{5, 10, 15}, 0},
		{"value greater than all thresholds", 20, []float64{5, 10, 15}, 3},
		{"empty thresholds slice", 10, []float64{}, 0},
		{"nil thresholds slice", 10, nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetPoint(tt.value, tt.thresholds)
			if actual != tt.expected {
				t.Errorf("GetPoint(%v, %v) = %v, want %v", tt.value, tt.thresholds, actual, tt.expected)
			}
		})
	}
}
