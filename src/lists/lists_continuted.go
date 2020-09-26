package lists

import (
	"errors"
)

// ModifiedLengthEncoding : Modify the result of problem 10 in such a way
// that if an element has no duplicates it is simply copied into the result list.
// Only elements with duplicates are transferred as (N E) lists.
// N/A in golang ?

// DecodeLength : decode an EncodedElement list to []string
func DecodeLength(lst []EncodedElement) []string {
	res := []string{}
	for _, e := range lst {
		for i := 0; i < e.occurences; i++ {
			res = append(res, e.element)
		}
	}
	return res
}

// Run-length encoding of a list (direct solution)

// DuplicateElements : Duplicate elements of a list
func DuplicateElements(lst []string) []string {
	duplicated := []string{}
	for i := 0; i < len(lst); i++ {
		duplicated = append(duplicated, lst[i], lst[i])
	}
	return duplicated
}

// ReplicateElements : Replicate the elements of a list a given number of times
func ReplicateElements(lst []string, occurence int) []string {
	replicated := []string{}
	for i := 0; i < len(lst); i++ {
		t := []string{}
		for j := 0; j < occurence; j++ {
			t = append(t, lst[i])
		}
		replicated = append(replicated, t...)
	}
	return replicated
}

// DropKthElement : Drop every K'th element from a list
func DropKthElement(lst []string, k int) []string {
	pruned := []string{}
	for index, e := range lst {
		if (index+1)%k != 0 {
			pruned = append(pruned, e)
		}
	}
	return pruned
}

// Split : Split a list into two parts; the length of the first part is given.
func Split(lst []string, l int) ([]string, []string, error) {
	if l > len(lst) {
		return []string{},
			[]string{},
			errors.New("Length given should be greater than given list length")
	}
	return lst[0:l], lst[l:len(lst)], nil
}

// ExtractSlice : Extract a slice from a list
func ExtractSlice(lst []string, s int, e int) ([]string, error) {
	if s > e {
		return []string{}, errors.New("Given start index should be lower then end index")
	}
	if e > len(lst) {
		return []string{}, errors.New("Given start index should be lower then end index")
	}
	if s <= 0 {
		return []string{}, errors.New("Given start index should be strictly greater than 0")
	}
	return lst[s-1 : e], nil
}

// Rotate : Rotate a list N places to the left. TODO
func Rotate(lst []string, n int) []string {
	r := make([]string, len(lst))
	for i, e := range lst {
		index := (i - n) % len(lst)
		if index < 0 {
			index = index + len(lst)
		}
		r[index] = e
	}
	return r
}

// RemoveKthElement : Remove the K'th element from a list
func RemoveKthElement(lst []string, k int) ([]string, error) {
	r := []string{}
	if k <= 0 {
		return []string{}, errors.New("k should be strictly greater than 0")
	}
	if k > len(lst) {
		return []string{}, errors.New("k should be lower than list len")
	}
	r = append(r, lst[0:k-1]...)
	r = append(r, lst[k:len(lst)]...)
	return r, nil
}
