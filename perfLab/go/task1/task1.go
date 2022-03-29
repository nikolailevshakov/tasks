package task1

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	// os.Args[1]
	Task1File = "./files/task1.txt"
)

func ReadFromFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func StringToStringSlice(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n")
}

func StringSliceToIntSlice(strSlice []string) []int {
	var intSlice []int
	for _, v := range strSlice {
		vTrimmed := strings.TrimSpace(v)
		vInt, err := strconv.Atoi(vTrimmed)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, vInt)
	}
	return intSlice
}

func Statistics(nums []int) (max, min, mean, perc90 int) {
	var sum int
	sort.Ints(nums)
	for _, v := range nums {
		sum += v
	}
	min = nums[0]
	max = nums[len(nums)-1]
	mean = sum / len(nums)
	perc90 = nums[int(float64(len(nums))*0.9)]
	return
}
