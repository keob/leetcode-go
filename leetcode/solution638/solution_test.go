package solution638

import "testing"

func TestShoppingOffers(t *testing.T) {
	tests := []struct {
		name    string
		price   []int
		special [][]int
		needs   []int
		want    int
	}{
		{
			"test 1",
			[]int{2, 5},
			[][]int{{3, 0, 5}, {1, 2, 10}},
			[]int{3, 2},
			14,
		},
		{
			"test 2",
			[]int{2, 3, 4},
			[][]int{{1, 1, 0, 4}, {2, 2, 1, 9}},
			[]int{1, 2, 1},
			11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shoppingOffers(tt.price, tt.special, tt.needs)
			if got != tt.want {
				t.Errorf("shoppingOffers() = %d, want %d", got, tt.want)
			}
		})
	}
}
