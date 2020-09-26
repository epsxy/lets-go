package lists

import (
	"errors"
)

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
func EliminateConsecutive(lst []string) []string {
	s := []string{}
	for _, e := range lst {
		if len(s) == 0 {
			s = append(s, e)
		}
		if s[len(s)-1] != e {
			s = append(s, e)
		}
	}
	return s
}

// PackConsecutive : pack consecutive elements of a list into sublists
func PackConsecutive(lst []string) [][]string {
	s := [][]string{}
	t := []string{}
	for _, e := range lst {
		if len(t) == 0 || t[len(t)-1] == e {
			t = append(t, e)
		} else {
			s = append(s, t)
			t = []string{}
			t = append(t, e)
		}
	}
	s = append(s, t)
	return s
}

// EncodedElement encoding structure
type EncodedElement struct {
	occurences int
	element    string
}

// EncodeLength : encode consecutive list elements
func EncodeLength(lst []string) []EncodedElement {
	p := PackConsecutive(lst)
	res := make([]EncodedElement, len(p))
	for i, e := range p {
		res[i] = EncodedElement{occurences: len(e), element: e[0]}
	}
	return res
}
