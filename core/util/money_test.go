package util_test

import (
	"testing"
	"transfer-api/core/util"

	"github.com/stretchr/testify/assert"
)

func TestFloatToCents(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{input: 1.234, expected: 123},
		{input: 1.235, expected: 124},
		{input: 123.456, expected: 12346},
		{input: 0.0, expected: 0},
		{input: -1.234, expected: -123},
		{input: -1.235, expected: -124},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := util.FloatToCents(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCentsToFloat64(t *testing.T) {
	tests := []struct {
		input    int
		expected float64
	}{
		{input: 123, expected: 1.23},
		{input: 124, expected: 1.24},
		{input: 12346, expected: 123.46},
		{input: 0, expected: 0.0},
		{input: -123, expected: -1.23},
		{input: -124, expected: -1.24},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := util.CentsToFloat64(test.input)
			assert.InDelta(t, test.expected, result, 0.0001)
		})
	}
}
