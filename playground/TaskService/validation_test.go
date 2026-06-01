package main

import "testing"

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected bool
	}{
		{
			name:     "valid title",
			title:    "title",
			expected: true,
		},
		{
			name:     "empty title",
			title:    "",
			expected: false,
		},
		{
			name:     "title with only spaces",
			title:    "    ",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validateTitle(test.title)

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestValidateArea(t *testing.T) {
	tests := []struct {
		name     string
		area     Area
		expected bool
	}{
		{
			name:     "valid area (backend)",
			area:     Backend,
			expected: true,
		},
		{
			name:     "invalid unknown area",
			area:     Unknown,
			expected: false,
		},
		{
			name:     "invalid negative area",
			area:     Area(-1),
			expected: false,
		},
		{
			name:     "invalid large area",
			area:     Area(1000),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validateArea(test.area)

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestValidatePriority(t *testing.T) {
	tests := []struct {
		name     string
		priority int
		expected bool
	}{
		{
			name:     "valid priority",
			priority: 5,
			expected: true,
		},
		{
			name:     "valid minimum priority",
			priority: 1,
			expected: true,
		},
		{
			name:     "valid maximum priority",
			priority: 10,
			expected: true,
		},
		{
			name:     "invalid zero priority",
			priority: 0,
			expected: false,
		},
		{
			name:     "invalid negative priority",
			priority: -1,
			expected: false,
		},
		{
			name:     "invalid large priority",
			priority: 100,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validatePriority(test.priority)

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestValidateEstimatedMinutes(t *testing.T) {
	tests := []struct {
		name             string
		estimatedMinutes int
		expected         bool
	}{
		{
			name:             "valid estimatedMinutes",
			estimatedMinutes: 60,
			expected:         true,
		},
		{
			name:             "valid minimum estimatedMinutes",
			estimatedMinutes: 1,
			expected:         true,
		},
		{
			name:             "invalid zero estimatedMinutes",
			estimatedMinutes: 0,
			expected:         false,
		},
		{
			name:             "invalid negative estimatedMinutes",
			estimatedMinutes: -1,
			expected:         false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validateEstimatedMinutes(test.estimatedMinutes)

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
