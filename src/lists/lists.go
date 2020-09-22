package lists

import "errors"

/*Generics are, as of today, not yet implemented in Go. Here we use int lists*/

/*Last : find the last element of an int list.*/
func Last(lst []int) (int, error) {
	if len(lst) == 0 {
		return 0, errors.New("Empty list provided")
	}
	return lst[len(lst)-1], nil
}

/*Antepenultimate : find the antepenultimate element of an int list*/
func Antepenultimate(lst []int) (int, error) {
	if len(lst) == 0 {
		return 0, errors.New("List size is 0, should be >=2")
	}
	if len(lst) == 1 {
		return 0, errors.New("List size is 1, should be >=2")
	}
	return lst[len(lst)-2], nil
}

/*ElementAt : find the kth element. k = 1 -> first element, k = 2 -> second element, etc.*/
func ElementAt(lst []int, k int) (int, error) {
	if k < 1 {
		return 0, errors.New("k-index should be >= 1, not < 1")
	}
	if k > len(lst) {
		return 0, errors.New("k-index is bigger than list length")
	}
	return lst[k-1], nil
}
