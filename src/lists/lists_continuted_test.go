package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*===================== DecodeLength Tests =====================*/
func TestShouldDecodeListLength(t *testing.T) {
	sample := []EncodedElement{
		EncodedElement{occurences: 3, element: "a"},
		EncodedElement{occurences: 1, element: "b"},
		EncodedElement{occurences: 2, element: "c"},
		EncodedElement{occurences: 5, element: "d"},
		EncodedElement{occurences: 2, element: "e"},
		EncodedElement{occurences: 2, element: "f"},
		EncodedElement{occurences: 1, element: "g"},
	}

	decoded := DecodeLength(sample)

	expected := []string{"a", "a", "a", "b", "c", "c", "d", "d", "d", "d", "d", "e", "e", "f", "f", "g"}

	assert.Equal(t, expected, decoded, "Should decode list length")
}

/*===================== DuplicateElements Tests =====================*/
func TestShouldDuplicateElements(t *testing.T) {
	sample := []string{"a", "b", "c"}

	duplicated := DuplicateElements(sample)
	expected := []string{"a", "a", "b", "b", "c", "c"}

	assert.Equal(t, expected, duplicated, "Should duplicated list elements")
}

/*===================== ReplicateElements Tests =====================*/
func TestShouldReplicateElements(t *testing.T) {
	sample := []string{"a", "b", "c"}

	replicated := ReplicateElements(sample, 3)
	expected := []string{"a", "a", "a", "b", "b", "b", "c", "c", "c"}

	assert.Equal(t, expected, replicated, "Should replicated list elements")
}

/*===================== DropKthElement Tests =====================*/
func TestShouldDropKthElement(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	pruned := DropKthElement(sample, 3)
	expected := []string{"a", "b", "d", "e", "g", "h"}

	assert.Equal(t, expected, pruned, "Should drop every 3 element")
}

/*===================== Split Tests =====================*/
func TestShouldReturnErrorWhileSplittingListWithGreaterSize(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	_, _, err := Split(sample, 9)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestShouldSplitList(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	first, second, err := Split(sample, 2)
	firstExpected := []string{"a", "b"}
	secondExpected := []string{"c", "d", "e", "f", "g", "h"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, firstExpected, first, "Should return correct first list")
	assert.Equal(t, secondExpected, second, "Should return correct second list")
}

func TestShouldSplitWithFirstEmptyList(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	first, second, err := Split(sample, 0)
	firstExpected := []string{}
	secondExpected := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, firstExpected, first, "Should return correct first list")
	assert.Equal(t, secondExpected, second, "Should return correct second list")
}
func TestShouldSplitWithSecondEmptyList(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	first, second, err := Split(sample, 8)
	firstExpected := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	secondExpected := []string{}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, firstExpected, first, "Should return correct first list")
	assert.Equal(t, secondExpected, second, "Should return correct second list")
}

/*===================== ExtractSlice Tests =====================*/
func TestShouldExtractSlice(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	sliced, err := ExtractSlice(sample, 2, 5)
	expected := []string{"b", "c", "d", "e"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, sliced, "Should slice list")
}

func TestShouldReturnErrorWhileExtractSliceWithWrongOrderingInBounds(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	_, err := ExtractSlice(sample, 5, 2)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestShouldReturnErrorWhileExtractSliceWithWrongGreaterBounds(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	_, err := ExtractSlice(sample, 5, 9)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestShouldReturnErrorWhileExtractSliceWithWrongLowerBounds(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	_, err := ExtractSlice(sample, 0, 9)

	assert.NotNil(t, err, "Error should not be nil")
}

/*===================== Rotate Tests =====================*/
// {"a", "b", "c", "d", "e", "f", "g", "h"}
//   0    1    2    3    4    5    6    7
//   4	  5    6    7    0    1    2    3
func TestRotateList(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	rotated := Rotate(sample, 4)
	expected := []string{"e", "f", "g", "h", "a", "b", "c", "d"}

	assert.Equal(t, expected, rotated, "Should rotate list")
}

func TestRotateLisWithNegativeInput(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	rotated := Rotate(sample, -2)
	expected := []string{"g", "h", "a", "b", "c", "d", "e", "f"}

	assert.Equal(t, expected, rotated, "Should rotate list with negative input")
}

/*===================== RemoveKthElement Tests =====================*/
func TestRemoveKthElement(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	removed, err := RemoveKthElement(sample, 3)
	expected := []string{"a", "b", "d", "e", "f", "g", "h"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, removed, "Should remove 3rd element")
}

func TestRemoveFirstElement(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	removed, err := RemoveKthElement(sample, 1)
	expected := []string{"b", "c", "d", "e", "f", "g", "h"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, removed, "Should remove first element")
}

func TestRemoveLastElement(t *testing.T) {
	sample := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	removed, err := RemoveKthElement(sample, 8)
	expected := []string{"a", "b", "c", "d", "e", "f", "g"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, removed, "Should remove last element")
}
