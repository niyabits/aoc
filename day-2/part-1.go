package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	partOne()
// }

func partOne() {
	content, _ := os.ReadFile("inp1.txt")
	reports := strings.Split(strings.TrimRight(string(content), "\n"), "\n")
	fmt.Println(len(reports))

	correct := make([]int, len(reports))

	for i, report := range reports {
		lis := strings.Split(report, " ")

		for j := 0; j < len(lis); j++ {
			if j != len(lis)-1 {
				a, _ := strconv.Atoi(lis[j])
				b, _ := strconv.Atoi(lis[j+1])

				res := math.Abs(float64(a - b))

				if res == 0 || res > 3 {
					correct[i] = 0
					break
				} else {
					correct[i] = 1
				}
			}
		}

		if !checkSorted(lis) {
			correct[i] = 0
		}
	}

	sum := 0
	for _, val := range correct {
		sum += val
	}

	fmt.Println(sum)
}

// func checkSorted(lisS []string) bool {
// 	lis := convertSliceToInt(lisS)
//
// 	asc := sort.SliceIsSorted(lis, func(i, j int) bool {
// 		return lis[i] < lis[j]
// 	})
//
// 	if asc == true {
// 		return true
// 	}
//
// 	desc := sort.SliceIsSorted(lis, func(i, j int) bool {
// 		return lis[i] > lis[j]
// 	})
//
// 	if desc == true {
// 		return true
// 	}
//
// 	return false
// }
//
// func convertSliceToInt(s []string) []int {
// 	var intList []int
// 	for _, str := range s {
// 		num, _ := strconv.Atoi(str)
// 		intList = append(intList, num)
// 	}
//
// 	return intList
// }
