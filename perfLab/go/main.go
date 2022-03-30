package main

import (
	"go/task1"
	"io/ioutil"
	"log"
)

const (
	task1File  = "./files/task1.txt"
	task2File1 = "./files/task2_1.txt"
	task2File2 = "./files/task2_2.txt"
	task3File1 = "./files/task3_1.txt"
)

func main() {
	// idealy code walk through the whole task3 directory
	// and take all .txt files in
	content := task1.ReadFromFile(task3File1)
	stringSlice := task1.StringToStringSlice(content)
	intSlice := task1.StringSliceToIntSlice(stringSlice)
}

func sumIntervals(intervals [][]int) []int {
	var sumSlice []int
	for _, row := range intervals {
		for i, val := range row {
			sumSlice[i] += val
		}
	}
	return sumSlice
}

func ReadFromFile(path []string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
