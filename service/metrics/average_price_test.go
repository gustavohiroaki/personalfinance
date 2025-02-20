package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateBasicAveragePrice(t *testing.T) {
	average := CalculateAveragePrice(20.0, 2.0)

	assert.Equal(t, 10.0, average)
}

func TestCalculateAveragePriceWithAnyValue0(t *testing.T) {
	average := CalculateAveragePrice(0.0, 2.0)

	assert.Equal(t, 0.0, average)
}
