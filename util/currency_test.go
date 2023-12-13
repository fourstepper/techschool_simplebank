package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrency(t *testing.T) {
	currency := RandomCurrency()
	result := IsSupportedCurrency(currency)
	require.True(t, result)

	falseCurrency := "IDR"
	result = IsSupportedCurrency(falseCurrency)
	require.False(t, result)
}
