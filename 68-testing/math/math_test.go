package math_test

import (
	"os"
	"testing"

	"github.com/james-computing/go-codeandlearn/68-testing/math"
)

var something string

func TestMain(m *testing.M) {
	something = "test"
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestMinInt(t *testing.T) {
	//f1(t)
	f2(t)
}

// Single test
func f1(t *testing.T) {
	result := math.MinInt(10, 5)
	if result != 5 {
		t.Errorf("expected: %d, got: %d", 5, result)
	}
}

// Table testing
func f2(t *testing.T) {
	t.Cleanup(func() {
		// Any cleanup required
	})

	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "100 < 1000",
			a:        100,
			b:        1000,
			expected: 100,
		},
		{
			name:     "-2 < -1",
			a:        -2,
			b:        -1,
			expected: -2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(something)
			result := math.MinInt(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	result := math.MaxInt(10, 100)
	if result != 100 {
		t.Errorf("expected: %d, got; %d", 100, result)
	}
}
