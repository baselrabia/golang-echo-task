package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReconstructRoute(t *testing.T) {
	testCases := []struct {
		name     string
		tickets  []Ticket
		expected []string
	}{
		{
			name: "Example input",
			tickets: []Ticket{
				{"LAX", "DXB"},
				{"JFK", "LAX"},
				{"SFO", "SJC"},
				{"DXB", "SFO"},
			},
			expected: []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
		},
		{
			name:     "Empty input",
			tickets:  []Ticket{},
			expected: []string{},
		},
		{
			name: "Single ticket",
			tickets: []Ticket{
				{"SFO", "LAX"},
			},
			expected: []string{"SFO", "LAX"},
		},
		{
			name: "Duplicate tickets",
			tickets: []Ticket{
				{"LAX", "DXB"},
				{"JFK", "LAX"},
				{"SFO", "SJC"},
				{"DXB", "SFO"},
				{"LAX", "DXB"},
			},
			expected: []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
		},
		{
			name: "Circular route",
			tickets: []Ticket{
				{"LAX", "DXB"},
				{"DXB", "SFO"},
				{"SFO", "LAX"},
			},
			expected: []string{"LAX", "DXB", "SFO", "LAX"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReconstructRoute(tc.tickets)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestReconstructRouteOnDisconnectedRoute(t *testing.T) {
	tickets := []Ticket{
		{"LAX", "DXB"},
		{"JFK", "LAX"},
		{"SFO", "SJC"},
		{"BOS", "MIA"},
	}
	actual := ReconstructRoute(tickets)

	expectedSegments := [][]string{
		{"JFK", "LAX", "DXB"},
		{"SFO", "SJC"},
		{"BOS", "MIA"},
	}

	for _, segment := range expectedSegments {
		assert.Subset(t, actual, segment, "actual itinerary does not contain expected segment")
	}
}
