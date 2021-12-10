package randomizers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCeil(t *testing.T) {
	t.Parallel()
	arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	res := findCeil(arr, 6, 0, 8)
	if arr[res] != 7 {
		t.Fatalf("Expected 7, got %d", res)
	}
}

func TestGetRandomWithDistribution(t *testing.T) {
	t.Parallel()
	arr := []int{1, 2, 9, 4}
	freq := []int{1, 2, 5, 6}
	res := GetRandomWithDistribution(arr, freq)
	assert.Contains(t, arr, res)

}

func TestRand75(t *testing.T) {
	t.Parallel()
	res := Rand75()
	assert.Contains(t, []int{0, 1}, res)
}
