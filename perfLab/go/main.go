package main

import (
	"go/task1"
	"strconv"
	"strings"
)

const (
	task1File  = "./files/task1.txt"
	task2File1 = "./files/task2_1.txt"
	task2File2 = "./files/task2_2.txt"
	task3File1 = "./files/task3_1.txt"
	task4File1 = "./files/task4.txt"
)

type Interval struct {
	start int
	end   int
	count int
}

func FromStringToIntervals(input string) []Interval {
	var intervals []Interval
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		str := strings.ReplaceAll(row, ":", "")
		strSlice := strings.Split(str, " ")
		if len(strSlice) < 2 {
			continue
		}
		start, _ := strconv.Atoi(strings.TrimSpace(strSlice[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(strSlice[1]))
		interval := Interval{
			start: start,
			end:   end,
			count: 1,
		}
		intervals = append(intervals, interval)
	}
	return intervals
}

func (interval1 Interval) isOverlepped(interval2 Interval) bool {
	if interval1.end <= interval2.start {
		return false
	}
	return true
}

func main() {
	input := task1.ReadFromFile(task4File1)
	intervals := FromStringToIntervals(input)
	// fmt.Println(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].isOverlepped(intervals[i+1]) {

		}
	}
}
