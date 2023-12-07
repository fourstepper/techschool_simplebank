package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	min       int64 = 0
	max       int64 = 1000
	lenString       = 6
)

func TestRandomInt(t *testing.T) {
	number := RandomInt(min, max)
	require.NotEmpty(t, number)
	require.Less(t, number, max)
}

func TestRandomString(t *testing.T) {
	lenString := 6
	word := RandomString(lenString)
	require.NotEmpty(t, word)
	require.Len(t, word, lenString)
}

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.NotEmpty(t, owner)
	require.Len(t, owner, lenString)
}

func TestRandomMoney(t *testing.T) {
	money := RandomMoney()
	require.NotEmpty(t, money)
	require.LessOrEqual(t, money, max)
}

func TestRandomCurrency(t *testing.T) {
	currencies := map[string]bool{
		"EUR": true,
		"USD": true,
		"CAD": true,
	}

	currency := RandomCurrency()
	exists := currencies[currency]
	require.NotEmpty(t, currency)
	require.True(t, exists)
}
