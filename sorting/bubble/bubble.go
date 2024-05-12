package bubble

func Sort(items []int) {
	for i := range len(items) - 1 {
		for j := range len(items) - 1 - i {
			if items[j] > items[j+1] {
				swap(items, j)
			}
		}
	}
}

func swap(items []int, index int) {
	tmp := items[index+1]
	items[index+1] = items[index]
	items[index] = tmp
}
