package lists

import (
	"errors"
	"math/rand"
)

//InsertElementAtPosition : Insert an element at a given position into a list.
func InsertElementAtPosition(lst []string, e string, p int) ([]string, error) {
	if p <= 0 {
		return []string{}, errors.New("position should be >= 1")
	}
	if p > len(lst)+1 {
		return []string{}, errors.New("position should be <= len(lst)")
	}
	r := []string{}
	r = append(r, lst[0:p-1]...)
	r = append(r, e)
	r = append(r, lst[p-1:len(lst)]...)
	return r, nil
}

//IntegersRange : Create a list containing all integers within a given range.
func IntegersRange(s int, e int) ([]int, error) {
	if s < 0 {
		return []int{}, errors.New("Start should be >= 0")
	}
	if s > e {
		return []int{}, errors.New("Start should be <= end")
	}
	r := []int{}
	for i := s; i <= e; i++ {
		r = append(r, i)
	}
	return r, nil
}

// RemoveKthElementInt : Remove Kth element from an int list
func RemoveKthElementInt(lst []int, k int) ([]int, error) {
	r := []int{}
	if k <= 0 {
		return []int{}, errors.New("k should be strictly greater than 0")
	}
	if k > len(lst) {
		return []int{}, errors.New("k should be lower than list len")
	}
	r = append(r, lst[0:k-1]...)
	r = append(r, lst[k:len(lst)]...)
	return r, nil
}

// ExtractRandoms : Extract a given number of randomly selected elements from a list.
func ExtractRandoms(lst []int, n int) ([]int, error) {
	r := []int{}
	t := append([]int{}, lst...)
	if len(lst) <= 0 {
		return []int{}, errors.New("List should not be empty")
	}
	if n > len(lst) {
		return []int{}, errors.New("Number should be below list length")
	}
	for index := 1; index <= n; {
		i := rand.Intn(len(t))
		r = append(r, t[i])
		// RemoveKthElement is a method from a previous problem
		t, _ = RemoveKthElementInt(t, i+1)
		index++
	}
	return r, nil
}

//Lotto : Draw N different random numbers from the set 1..M.
func Lotto(n int, m int) ([]int, error) {
	if m <= 0 {
		return []int{}, errors.New("m should be >= 1")
	}
	if n > m {
		return []int{}, errors.New("n should be <= m")
	}
	s, _ := IntegersRange(1, m)
	r, err := ExtractRandoms(s, n)
	if err != nil {
		return []int{}, errors.New("ExtractRandoms failed")
	}
	return r, nil
}

//Problem 25
//Generate a random permutation of the elements of a list.

//Problem 26
//(**) Generate the combinations of K distinct objects chosen from the N elements of a list

//Problem 27
//Group the elements of a set into disjoint subsets.

//Problem 28
//Sorting a list of lists according to length of sublists
