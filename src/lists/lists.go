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
