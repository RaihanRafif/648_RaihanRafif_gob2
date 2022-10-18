package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFailTheFunction(t *testing.T) {
	result := TheFunction(2)

	require.Equal(t, false, result, "Result has to be false")
	fmt.Println("TestFailFunction Eksekusi Terhenti")
}

func TestTheFunction(t *testing.T) {
	result := TheFunction(2)

	assert.Equal(t, true, result, "Result has to be true")
	fmt.Println("TestFunction Eksekusi Terhenti")

}
