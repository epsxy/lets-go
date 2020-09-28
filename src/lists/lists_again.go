package lists

import "errors"

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

//Problem 23
//Extract a given number of randomly selected elements from a list.

//Problem 24
//Lotto: Draw N different random numbers from the set 1..M.

//Problem 25
//Generate a random permutation of the elements of a list.

//Problem 26
//(**) Generate the combinations of K distinct objects chosen from the N elements of a list

//Problem 27
//Group the elements of a set into disjoint subsets.

//Problem 28
//Sorting a list of lists according to length of sublists
