package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expectec %v", got, expected)
	}
}
