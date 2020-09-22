package arithmetics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*===================== IsPrime Tests =====================*/
func TestVerifyNumberIsNotPrime(t *testing.T) {
	res := IsPrime(4)
	assert.Equal(t, res, false, "Should not be prime")
}
func TestVerifyNumberIsPrime(t *testing.T) {
	res := IsPrime(23)
	assert.Equal(t, res, true, "Should be prime")
}
