package bubble

// this program shows the implementation of the bubble sort and calling that with goroutines
// sorting of a large array is divied into 4 parts and each part is sorted by a goroutine
// then the sorted parts are merged to get the final sorted array
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func RunConcurrentSorts() {
	fmt.Println("Input a series of integers")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	userInStr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	userInInts := make([]int, 0, len(userInStr))
	for _, v := range userInStr {
		intV, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid input %s cannot convert to integer\n", v)
			continue
		}
		userInInts = append(userInInts, intV)
	}

	fmt.Println(userInInts)

	var partitions = partitionSlice(userInInts, 4)
	wg := &sync.WaitGroup{}
	go sortList(partitions[0], wg) // go routine 1
	go sortList(partitions[1], wg) // go routine 2
	go sortList(partitions[2], wg) // go routine 3
	go sortList(partitions[3], wg) // go routine 4
	wg.Wait()
	finalList1 := sortPartitions(partitions[0], partitions[1])
	finalList2 := sortPartitions(partitions[2], partitions[3])
	final := sortPartitions(finalList1, finalList2)
	fmt.Println(final)
}

func partitionSlice(originalList []int, nPartitions int) [][]int {
	var partitionSize int = len(originalList) / nPartitions
	var remainder int = len(originalList) % nPartitions

	var returnLists [][]int = make([][]int, 0, nPartitions)

	var prevUpperLimit = 0
	for i := 0; i < nPartitions; i++ {
		var partition []int
		if remainder > 0 {
			partition = originalList[prevUpperLimit : prevUpperLimit+partitionSize+1]
			prevUpperLimit += partitionSize + 1
			remainder--
		} else {
			partition = originalList[prevUpperLimit : prevUpperLimit+partitionSize]
			prevUpperLimit += partitionSize
		}
		returnLists = append(returnLists, partition)
	}
	return returnLists
}

func sortPartitions(part1, part2 []int) []int {
	pos1, pos2 := 0, 0
	var returnSlice = make([]int, 0, len(part1)+len(part2))
	for {
		if pos1 == len(part1) && pos2 != len(part2) {
			returnSlice = append(returnSlice, part2[pos2:]...)
			break
		}
		if pos2 == len(part2) && pos1 != len(part1) {
			returnSlice = append(returnSlice, part1[pos1:]...)
			break
		}
		if part1[pos1] < part2[pos2] {
			returnSlice = append(returnSlice, part1[pos1])
			pos1++
		} else {
			returnSlice = append(returnSlice, part2[pos2])
			pos2++
		}
	}
	return returnSlice
}

// based on bubble sort with a wait group
func sortList(list []int, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Before sorting: ", list)
	for i := range len(list) - 1 {
		for j := range len(list) - 1 - i {
			if list[j] > list[j+1] {
				swap(list, j)
			}
		}
	}
	fmt.Println("After sorting: ", list)
	wg.Done()
}
