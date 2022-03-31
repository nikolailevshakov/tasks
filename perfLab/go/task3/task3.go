package task3

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// task 3 - Put in main
// var floatContent [][]float64
// content := ReadFromFiles("./files")
// for _, contentFile := range content {
//   stringSlice := task1.StringToStringSlice(contentFile)
//   floatSlice := StringSliceToFloatSlice(stringSlice)
//   floatContent = append(floatContent, floatSlice)
// }
// sumSlice := sumIntervals(floatContent)
// maxIndex := findMax(sumSlice)
// fmt.Println(maxIndex)

func sumIntervals(intervals [][]float64) []float64 {
	sumSlice := make([]float64, 16)
	for _, row := range intervals {
		for i, val := range row {
			sumSlice[i] += val
		}
	}
	return sumSlice
}

func findMax(slice []float64) int {
	var max float64
	var maxIndex int
	for index, val := range slice {
		if val > max {
			max, maxIndex = val, index
		}
	}
	return maxIndex
}

func ReadFromFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func ReadFromFiles(directory string) []string {
	var allContents []string
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "task3") {
			allContents = append(allContents, ReadFromFile(directory+"/"+file.Name()))
		}
	}
	return allContents
}

func StringSliceToFloatSlice(strSlice []string) []float64 {
	var floatSlice []float64
	for _, v := range strSlice {
		vTrimmed := strings.TrimSpace(v)
		vFloat, err := strconv.ParseFloat(vTrimmed, 64)
		if err != nil {
			log.Fatal(err)
		}
		floatSlice = append(floatSlice, vFloat)
	}
	return floatSlice
}
