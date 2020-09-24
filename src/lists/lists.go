package lists

import "errors"

// NB:
// Generics are, as of today, not yet implemented in Go. Here we use int lists

// Last : find the last element of an int list
func Last(lst []int) (int, error) {
	if len(lst) == 0 {
		return 0, errors.New("Empty list provided")
	}
	return lst[len(lst)-1], nil
}

// Antepenultimate : find the antepenultimate element of an int list
func Antepenultimate(lst []int) (int, error) {
	if len(lst) == 0 {
		return 0, errors.New("List size is 0, should be >=2")
	}
	if len(lst) == 1 {
		return 0, errors.New("List size is 1, should be >=2")
	}
	return lst[len(lst)-2], nil
}

// ElementAt : find the kth element of a list
func ElementAt(lst []int, k int) (int, error) {
	if k < 1 {
		return 0, errors.New("k-index should be >= 1, not < 1")
	}
	if k > len(lst) {
		return 0, errors.New("k-index is bigger than list length")
	}
	return lst[k-1], nil
}

// Length : compute the length of a list
func Length(lst []int) int {
	return len(lst)
}

// Reverse : reverse a list
func Reverse(lst []int) []int {
	reversedList := make([]int, len(lst))
	for i, e := range lst {
		reversedList[len(lst)-1-i] = e
	}
	return reversedList
}

// IsPalindrome : determinate if a given list is a palindrome
func IsPalindrome(lst []int) bool {
	for i := 0; i < len(lst)/2; i++ {
		if lst[i] != lst[len(lst)-1-i] {
			return false
		}
	}
	return true
}

// NestedString : describe a nested string list structure
// e.g. [ 1 [1.1 1.2] 2 [2.1 2.2 [2.2.1] 2.3] 3 [3.1] ]
type NestedString struct {
	next    []*NestedString
	element string
}

// Flatten : flatten a nested list structure
func Flatten(root []*NestedString) []string {
	s := []string{}
	for _, child := range root {
		s = append(s, child.element)
		s = append(s, Flatten(child.next)...)
	}
	return s
}

// EliminateConsecutive : eliminate consecutive elements of a list

// PackConsecutive : pack consecutive elements of a list

// EncodeLength : encode consecutive list elements
