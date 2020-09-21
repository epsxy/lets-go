package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastElementOfEmptyList(t *testing.T) {
	lst := make([]int, 0)
	_, err := Last(lst)
	if err == nil {
		t.Errorf("Last([]) err nil want 'Empty list provided' ")
	}
	assert.EqualError(t, err, "Empty list provided", "Should return an error")
}

func TestGetLastElementOfOneElementList(t *testing.T) {
	lst := make([]int, 1)
	lst[0] = 1
	last, err := Last(lst)
	if err != nil {
		t.Errorf("Error is not nil")
	}
	if last != 1 {
		t.Errorf("Last([1 2 3]) = %d wants %d", last, 1)
	}
}

func TestGetLastElementOfBigList(t *testing.T) {
	lst := make([]int, 3)
	lst[0] = 1
	lst[1] = 2
	lst[2] = 3
	last, err := Last(lst)
	if err != nil {
		t.Errorf("Error is not nil")
	}
	if last != 3 {
		t.Errorf("Last([1 2 3]) = %d wants %d", last, 3)
	}
}
