package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBigger(t *testing.T){
	res01 := isBigger(3, 4)
	require.True(t, true, res01)

	res02 := isBigger (4, 3)
	require.False(t, false, res02)
}

