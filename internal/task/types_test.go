package task

import "testing"

func TestAreaString(t *testing.T) {
	tests := []struct {
		name     string
		area     Area
		expected string
	}{
		{
			name:     "backend area",
			area:     Backend,
			expected: "Backend",
		},
		{
			name:     "english area",
			area:     English,
			expected: "English",
		},
		{
			name:     "guitar area",
			area:     Guitar,
			expected: "Guitar",
		},
		{
			name:     "algorithms area",
			area:     Algorithms,
			expected: "Algorithms",
		},
		{
			name:     "university area",
			area:     University,
			expected: "University",
		},
		{
			name:     "invalid area",
			area:     Area(100),
			expected: "Unknown Area",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.area.String()

			if result != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, result)
			}
		})
	}
}

func TestAreaIsValid(t *testing.T) {
	tests := []struct {
		name     string
		area     Area
		expected bool
	}{
		{
			name:     "backend area",
			area:     Backend,
			expected: true,
		},
		{
			name:     "english area",
			area:     English,
			expected: true,
		},
		{
			name:     "guitar area",
			area:     Guitar,
			expected: true,
		},
		{
			name:     "algorithms area",
			area:     Algorithms,
			expected: true,
		},
		{
			name:     "university area",
			area:     University,
			expected: true,
		},
		{
			name:     "invalid large area",
			area:     Area(100),
			expected: false,
		},
		{
			name:     "invalid negative area",
			area:     Area(-1),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.area.IsValid()

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestStatusString(t *testing.T) {
	tests := []struct {
		name     string
		status   Status
		expected string
	}{
		{
			name:     "Todo status",
			status:   Todo,
			expected: "To Do",
		},
		{
			name:     "complete status",
			status:   Done,
			expected: "Done",
		},
		{
			name:     "In_progress status",
			status:   In_progress,
			expected: "In Progress",
		},
		{
			name:     "Blocked status",
			status:   Blocked,
			expected: "Blocked",
		},
		{
			name:     "Cancelled status",
			status:   Cancelled,
			expected: "Cancelled",
		},
		{
			name:     "invalid large status",
			status:   Status(100),
			expected: "Unknown Status",
		},
		{
			name:     "invalid negative status",
			status:   Status(-1),
			expected: "Unknown Status",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.status.String()

			if result != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, result)
			}
		})
	}
}

func TestAStatusIsValid(t *testing.T) {
	tests := []struct {
		name     string
		status   Status
		expected bool
	}{
		{
			name:     "Todo status",
			status:   Todo,
			expected: true,
		},
		{
			name:     "complete status",
			status:   Done,
			expected: true,
		},
		{
			name:     "In_progress status",
			status:   In_progress,
			expected: true,
		},
		{
			name:     "Blocked status",
			status:   Blocked,
			expected: true,
		},
		{
			name:     "Cancelled status",
			status:   Cancelled,
			expected: true,
		},
		{
			name:     "invalid large status",
			status:   Status(100),
			expected: false,
		},
		{
			name:     "invalid negative status",
			status:   Status(-1),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.status.IsValid()

			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
