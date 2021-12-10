package array_test

import (
	"testing"

	"github.com/Ahmed-Sermani/algos/randomized/array"
)

func TestQuickSelect(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name string
		data []int
		k    int
		want int
	}{
		{
			name: "single",
			data: []int{1},
			k:    0,
			want: 1,
		},
		{
			name: "two",
			data: []int{1, 2},
			k:    0,
			want: 1,
		},
		{
			name: "three",
			data: []int{1, 2, 3},
			k:    0,
			want: 1,
		},
		{
			name: "three",
			data: []int{1, 2, 3},
			k:    1,
			want: 2,
		},
		{
			name: "three",
			data: []int{1, 2, 3},
			k:    2,
			want: 3,
		},
		{
			name: "unsorted",
			data: []int{3, 2, 1},
			k:    0,
			want: 1,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := array.QuickSelect(test.data, test.k)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
