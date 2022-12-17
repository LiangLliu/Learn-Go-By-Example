package arrays

import "testing"

func TestSum(t *testing.T) {
	/*
		init array

		[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
		[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}
	*/

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
