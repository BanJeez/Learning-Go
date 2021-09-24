package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result int
	/* The line above initialises result as an empty int variable.
	The for loop below tells it to apply func f to every element and if it
	returns true, i.e. is numeric, then to add each element that is true
	 to the empty variable result and if not to ignore the whole argument.
	 The last line tells it to return the result.   */

	for _, element := range tab {
		if f(element) {
			result++
		}
	}
	return result
}
