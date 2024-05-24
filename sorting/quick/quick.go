package quick

func Sort(items []int) {
	// take a random pivot point
	pivot := getPivot(items)
	// swap the pivot point with the last element of the array
	swap(items, pivot, len(items)-1)
	// make the pivot elements placed in the correct position
	// take i, j iterators from the left side of the array i=-1, j=0
	i, j := -1, 0
	// move j to the right until we find an element less than or equal the pivot
	for j < len(items)-1 {
		// if we find an element less than or equal the pivot, we increment i and swap the elements at i and j
		if items[j] <= items[len(items)-1] {
			i++
			swap(items, i, j)
		}
		j++
	}
	// finally, swap the pivot element with the element at i+1
	i++
	swap(items, i, len(items)-1)
	// now the pivot element is placed in the correct position
	// take the left and right sub-arrays and repeat the process recursively
	if i > 0 {
		Sort(items[:i])
	}
	if i < len(items)-1 {
		Sort(items[i+1:])
	}
}

func getPivot(items []int) int {
	// there may be several ways to choose a pivot point
	// we can choose the first element, the last element, the middle element, or a random element
	// here we choose the last element as the pivot
	return len(items) - 1
}

func swap(items []int, i, j int) {
	tmp := items[i]
	items[i] = items[j]
	items[j] = tmp
}
