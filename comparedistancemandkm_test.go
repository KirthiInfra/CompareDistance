package compareDistance

import (
	"testing"
)

func TestCompareDistance(t *testing.T) {
	tests := []struct {
		name     string
		d1       distance
		d2       distance
		expected bool
	}{
		{
			name:     "1000 meters equal to 1000 meters",
			d1:       distance{value: 1000, unit: m},
			d2:       distance{value: 1000, unit: m},
			expected: true,
		},
		{
			name:     "1 meter does not equal to 2 meters",
			d1:       distance{value: 1, unit: m},
			d2:       distance{value: 2, unit: m},
			expected: false,
		},
		{
			name:     "1000 meters equal to 1 kilometer",
			d1:       distance{value: 1000, unit: m},
			d2:       distance{value: 1, unit: km},
			expected: true,
		},
		{
			name:     "1 kilometer equals to 1000 meters",
			d1:       distance{value: 1, unit: km},
			d2:       distance{value: 1000, unit: m},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d1.IsDistanceEqual(tt.d2)
			if got != tt.expected {
				t.Errorf("CompareDistances() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCreateDistanceStruct_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value int
		unit  unit
	}{
		{"Valid meters", 1000, m},
		{"Valid kilometers", 2, km},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateDistancesStruct(tt.value, tt.unit)
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}

func TestCreateDistanceStruct_NegativeValue(t *testing.T) {
	_, err := CreateDistancesStruct(-1, m)
	if err == nil {
		t.Errorf("Expected error for negative value, got none")
	}
}

func TestCreateDistanceStruct_InvalidUnit(t *testing.T) {
	_, err := CreateDistancesStruct(1, "kmm")
	if err == nil {
		t.Errorf("Expected error for invalid unit, got none")
	}
}
