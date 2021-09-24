package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	/* This line above initialises the result variable as 0. The for loop below iterates over
	   the slice of int that a represents and adds each element to the empty variable result. */
	for _, nbr := range a {
		result = append(result, f(nbr))
	}

	/*The loop above loops over every element in range a, applies the f function to each
	element and then appends to the empty bool slice created with var result []bool above it */

	/* The variable boolA, which uses the make method, creates a slice of type bool and
	then uses the result variable as the size of the slice which matches to the slice of
	int represented by a. */

	/*The for loop below loops over the slice of int a again but this time says to apply
	function referred to by Map on each element of a and then pass it to the slice
	represented by boolA */

	/*The return function below returns what is is in the slice of bool represented by boolA
	  which is the result of the f function applied to each element in the slice of int
	  represented by a. */
	return result
}
