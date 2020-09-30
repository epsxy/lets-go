package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*===================== InsertElementAtPosition Tests =====================*/
func TestShouldInsertElementInList(t *testing.T) {
	sample := []string{"a", "b", "c", "d"}

	result, err := InsertElementAtPosition(sample, "z", 2)
	expected := []string{"a", "z", "b", "c", "d"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result, "Should insert element at index 2")
}
func TestShouldInsertElementAtBeginningOfList(t *testing.T) {
	sample := []string{"a", "b", "c", "d"}

	result, err := InsertElementAtPosition(sample, "z", 1)
	expected := []string{"z", "a", "b", "c", "d"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result, "Should insert element at first index")
}
func TestShouldInsertElementAtTheEndOfList(t *testing.T) {
	sample := []string{"a", "b", "c", "d"}

	result, err := InsertElementAtPosition(sample, "z", 5)
	expected := []string{"a", "b", "c", "d", "z"}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result, "Should insert element at last index")
}

func TestShouldFailIfPositionStrictlyLowerThanOne(t *testing.T) {
	sample := []string{"a", "b", "c", "d"}

	_, err := InsertElementAtPosition(sample, "z", 0)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestShouldFailIfPositionStrictlyGreaterThanLengthPlusOne(t *testing.T) {
	sample := []string{"a", "b", "c", "d"}

	_, err := InsertElementAtPosition(sample, "z", 6)

	assert.NotNil(t, err, "Error should not be nil")
}

/*===================== IntegersRange Tests =====================*/
func TestShouldGenerateIntegersRange(t *testing.T) {

	result, err := IntegersRange(4, 10)
	expected := []int{4, 5, 6, 7, 8, 9, 10}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, result, "Should generate range from 4 to 10")
}

func TestShouldGenerateIntegersRangeWithLowerBoundError(t *testing.T) {

	_, err := IntegersRange(-1, 10)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestShouldGenerateIntegersRangeWithBoundsOrderError(t *testing.T) {

	_, err := IntegersRange(4, 2)

	assert.NotNil(t, err, "Error should not be nil")
}

/*===================== RemoveKthElementInt Tests =====================*/
func TestRemoveKthElementInt(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8}

	removed, err := RemoveKthElementInt(sample, 3)
	expected := []int{1, 2, 4, 5, 6, 7, 8}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, removed, "Should remove 3rd element")
}

func TestRemoveFirstElementInt(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8}

	removed, err := RemoveKthElementInt(sample, 1)
	expected := []int{2, 3, 4, 5, 6, 7, 8}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, removed, "Should remove first element")
}

func TestRemoveLastElementInt(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8}

	removed, err := RemoveKthElementInt(sample, 8)
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, expected, removed, "Should remove last element")
}

func TestRemoveLastElementIntShouldReturnErrorWithLowerBound(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8}

	_, err := RemoveKthElementInt(sample, 0)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestRemoveLastElementIntShouldReturnErrorWithGreaterBound(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8}

	_, err := RemoveKthElementInt(sample, 9)

	assert.NotNil(t, err, "Error should not be nil")
}

/*===================== ExtractRandoms Tests =====================*/
func TestShouldExtractRandoms(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := ExtractRandoms(sample, 4)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, []int{2, 8, 10, 1}, res, "Should extract randoms")
}

/*===================== Lotto Tests =====================*/
func TestLottoShouldWork(t *testing.T) {
	res, err := Lotto(5, 10)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, []int{2, 8, 3, 9, 7}, res, "Should do lotto")
}

func TestLottoShouldFailIfMBelowOne(t *testing.T) {
	_, err := Lotto(0, 0)

	assert.NotNil(t, err, "Error should not be nil")
	assert.EqualError(t, err, "m should be >= 1")
}

func TestLottoShouldFailIfNGreaterThanM(t *testing.T) {
	_, err := Lotto(11, 10)

	assert.NotNil(t, err, "Error should not be nil")
	assert.EqualError(t, err, "n should be <= m")
}

func TestLottoShouldFailIfExtractRandomsFails(t *testing.T) {
	res, err := Lotto(1, 1)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, []int{1}, res, "Should do very small lotto")
}
