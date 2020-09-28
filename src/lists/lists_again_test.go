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
