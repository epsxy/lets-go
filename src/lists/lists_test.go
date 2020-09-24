package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*===================== Last Tests =====================*/
func TestGetLastElementOfEmptyList(t *testing.T) {
	lst := []int{}
	_, err := Last(lst)
	if err == nil {
		t.Errorf("Last([]) err nil want 'Empty list provided' ")
	}
	assert.EqualError(t, err, "Empty list provided", "Should return an error")
}

func TestGetLastElementOfOneElementList(t *testing.T) {
	lst := []int{1}
	last, err := Last(lst)
	if err != nil {
		t.Errorf("Error is not nil")
	}
	if last != 1 {
		t.Errorf("Last([1 2 3]) = %d wants %d", last, 1)
	}
}

func TestGetLastElementOfBigList(t *testing.T) {
	lst := []int{1, 2, 3}
	last, err := Last(lst)
	if err != nil {
		t.Errorf("Error is not nil")
	}
	if last != 3 {
		t.Errorf("Last([1 2 3]) = %d wants %d", last, 3)
	}
}

/*===================== Antepenultimate Tests =====================*/
func TestGetAntepenultimateElementOfEmptyList(t *testing.T) {
	lst := []int{}
	_, err := Antepenultimate(lst)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
	assert.EqualError(t, err, "List size is 0, should be >=2", "Should return an error")
}

func TestGetAntepenultimateElementOfOneElementList(t *testing.T) {
	lst := []int{1}
	_, err := Antepenultimate(lst)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
	assert.EqualError(t, err, "List size is 1, should be >=2", "Should return an error")
}

func TestGetAntepenultimateElementOfBigList(t *testing.T) {
	lst := []int{1, 2, 3}
	antepenultimate, err := Antepenultimate(lst)
	if err != nil {
		t.Errorf("Error is not nil")
	}
	if antepenultimate != 2 {
		t.Errorf("Antepenultimate([1 2 3]) = %d wants %d", antepenultimate, 2)
	}
}

/*===================== ElementAt Tests =====================*/
func TestGetKthElementOfListWithIncorrectKValue(t *testing.T) {
	lst := []int{}
	_, err := ElementAt(lst, 0)
	assert.EqualError(t, err, "k-index should be >= 1, not < 1", "Should return an error")
}
func TestGetKthElementOfEmptyList(t *testing.T) {
	lst := []int{}
	_, err := ElementAt(lst, 2)
	assert.EqualError(t, err, "k-index is bigger than list length", "Should return an error")
}

func TestGetKthElementOfBigList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	kthElement, err := ElementAt(lst, 4)
	if err != nil {
		t.Errorf("Error is not nil")
	}
	if kthElement != 4 {
		t.Errorf("Antepenultimate([1 2 3]) = %d wants %d", kthElement, 4)
	}
}

/*===================== Length Tests =====================*/
func TestShouldGetListLength(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}

	l := Length(lst)

	assert.Equal(t, l, 5, "List size should be 5")
}

/*===================== Reverse Tests =====================*/
func TestShouldReverseEmptyList(t *testing.T) {
	lst := []int{}

	r := Reverse(lst)

	assert.Equal(t, r, []int{}, "Should be empty list")
}

func TestShouldReverseList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}

	r := Reverse(lst)

	assert.Equal(t, r, []int{5, 4, 3, 2, 1}, "Should reverse list")
}

/*===================== IsPalindrome Tests =====================*/
func TestShouldDeterminateNotPalindromeEvenList(t *testing.T) {
	lst := []int{1, 2, 2, 4}

	isPal := IsPalindrome(lst)

	assert.False(t, isPal)
}
func TestShouldDeterminateNotPalindromeOddList(t *testing.T) {
	lst := []int{1, 2, 3, 2, 4}

	isPal := IsPalindrome(lst)

	assert.False(t, isPal)
}

func TestShouldDeterminatePalindromeEmptyList(t *testing.T) {
	lst := []int{}

	isPal := IsPalindrome(lst)

	assert.True(t, isPal)
}

func TestShouldDeterminatePalindromeEvenList(t *testing.T) {
	lst := []int{1, 2, 3, 3, 2, 1}

	isPal := IsPalindrome(lst)

	assert.True(t, isPal)
}
func TestShouldDeterminatePalindromeOddList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 3, 2, 1}

	isPal := IsPalindrome(lst)

	assert.True(t, isPal)
}

/*===================== Flatten Tests =====================*/
func TestShouldFlattenList(t *testing.T) {
	nested := []*NestedString{
		&NestedString{[]*NestedString{
			&NestedString{[]*NestedString{}, "1.1"},
			&NestedString{[]*NestedString{}, "1.2"},
		}, "1"},
		&NestedString{[]*NestedString{
			&NestedString{[]*NestedString{}, "2.1"},
			&NestedString{[]*NestedString{
				&NestedString{[]*NestedString{}, "2.2.1"},
			}, "2.2"},
			&NestedString{[]*NestedString{}, "2.3"},
		}, "2"},
		&NestedString{[]*NestedString{
			&NestedString{[]*NestedString{}, "3.1"},
		}, "3"},
	}

	flattened := Flatten(nested)
	expected := []string{"1", "1.1", "1.2", "2", "2.1", "2.2", "2.2.1", "2.3", "3", "3.1"}

	assert.Equal(t, expected, flattened, "Should flatten nested list")
}
