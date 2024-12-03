package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}

func partTwo() {
	content, _ := os.ReadFile("inp1.txt")
	reports := strings.Split(strings.TrimRight(string(content), "\n"), "\n")

	var correct int
	var dampedReports int
	for _, report := range reports {
		r := strToIntArr(strings.Split(report, " "))
		status, _ := checkReport(r)

		if status == true {
			correct += 1
		}

		if status == false {
			// remove the bad value from the slice
			for i := 0; i < len(r); i++ {
				if i == len(r)-1 {
					r = r[:i]
				} else {
					r = append(r[:i], r[i+1:]...)
				}
				newStatus, _ := checkReport(r)

				if newStatus == true {
					dampedReports += 1
				}
			}
		}
	}

	fmt.Println(correct)
	fmt.Println(dampedReports)
	fmt.Println(correct + dampedReports)
}

func checkReport(report []int) (bool, int) {
	var isIncreasing bool

	for i := range report {
		// avoid out of index for i+1
		if i < len(report)-1 {
			currReport := report[i]
			nextReport := report[i+1]

			// 1.
			if currReport > nextReport {
				if i != 0 && isIncreasing == true {
					return false, i
				}

				isIncreasing = false

				// 2.
				if currReport-nextReport > 3 {
					return false, i
				}
			} else if currReport < nextReport {
				if i != 0 && isIncreasing == false {
					return false, i
				}

				isIncreasing = true

				if nextReport-currReport > 3 {
					return false, i
				}
			} else {
				// both of them are equal
				return false, i
			}
		}

	}

	return true, 0
}

func strToIntArr(list []string) []int {
	intArr := make([]int, len(list))
	for i, str := range list {
		num, _ := strconv.Atoi(str)
		intArr[i] = num
	}

	return intArr
}
