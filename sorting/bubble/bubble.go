package bubble

func Sort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-1-i; j++ {
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
