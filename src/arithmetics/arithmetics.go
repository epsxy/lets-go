package arithmetics

import (
	"math/big"
)

/*IsPrime : determinate if a number is prime or not*/
func IsPrime(number int64) bool {
	return big.NewInt(number).ProbablyPrime(0)
}
