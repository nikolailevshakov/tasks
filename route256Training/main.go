package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input := readStdInputSlice()
	input := []string{"3", "3", "cba", "aaz", "2", "bc", "ab", "1", "a", "a"}
	passwordChecks := originalPassword(input)
	writeStdOutputSlice(passwordChecks)
}

// Задание 4. Тренировочный раунд.
func originalPassword(inputData []string) []string {
	var answer []string
	numberOfPasswords, _ := strconv.Atoi(inputData[0])
	for i := 0; i < numberOfPasswords*3; i = i + 3 {
		if compareStrings(inputData[i+2], inputData[i+3]) {
			answer = append(answer, "YES")
		} else {
			answer = append(answer, "NO")
		}
	}
	return answer
}

func compareStrings(firstHalf, secondHalf string) bool {
	for _, firstHalfLetter := range firstHalf {
		for _, secondHalfLetter := range secondHalf {
			if firstHalfLetter < secondHalfLetter {
				return true
			}
		}
	}
	return false
}

// Задание 3. Тренировочный раунд. Слишком медленно.
func happyOrder(orderNum string) string {
	orderNumInt, _ := strconv.Atoi(orderNum)
	for i := orderNumInt; ; i++ {
		if isPalindrome(i) {
			return strconv.Itoa(i - orderNumInt)
		}
	}
	return "No palindrom!"
}

func isPalindromeStack(num int) bool {
	numStr := strconv.Itoa(num)
	var reverseNumStr string
	for i := len(numStr) - 1; i >= 0; i-- {
		reverseNumStr += string(numStr[i])
		if string(numStr[i]) != string(reverseNumStr[len(numStr)-1-i]) {
			return false
		}
	}
	if numStr == reverseNumStr {
		return true
	}
	return false
}

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	var reversed, remainder, original int
	original = num
	for num != 0 {
		remainder = num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	return original == reversed
}

// Задание 2. Тренировочный раунд.
//TODO: Причесать
func forgottenPincode(numbers string) []int {
	var pinOptionsDigits []string = strings.Split(numbers, "")
	var pincodeCombinations []int

	sortedPinOptionsInts := bubbleSortSlice(sliceStringsToSliceInts(pinOptionsDigits))
	pincodeCombinations = sliceSlicesToSliceInts(permutations(sortedPinOptionsInts))

	return bubbleSortSlice(pincodeCombinations)
}

func bubbleSortSlice(intsPincode []int) []int {
	for j := 0; j < len(intsPincode)-1; j++ {
		for i := 0; i < len(intsPincode)-1; i++ {
			if intsPincode[i] > intsPincode[i+1] {
				intsPincode[i], intsPincode[i+1] = intsPincode[i+1], intsPincode[i]
			}
		}
	}
	return intsPincode
}

func sliceStringsToSliceInts(stringsSlice []string) []int {
	var ints []int
	for _, v := range stringsSlice {
		num, _ := strconv.Atoi(v)
		ints = append(ints, num)
	}
	return ints
}

func sliceIntsToSliceStrings(intSlice []int) []string {
	var stringsSlice []string
	for _, v := range intSlice {
		str := strconv.Itoa(v)
		stringsSlice = append(stringsSlice, str)
	}
	return stringsSlice
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func sliceItoa(ints []int, num1, num2, num3 int) string {
	return strconv.Itoa(ints[num1]) + strconv.Itoa(ints[num2]) + strconv.Itoa(ints[num3])
}

func sliceSlicesToSliceInts(sliceSlices [][]int) []int {
	var result []int
	for _, slice := range sliceSlices {
		var numStr string
		for _, d := range slice {
			numStr += strconv.Itoa(d)
		}
		num, _ := strconv.Atoi(numStr)
		result = append(result, num)
	}
	return result
}

// Задание 1. Тренировочный раунд.
func operatorsCode(phoneNumber string) string {
	return phoneNumber[1:4]
}

// Функции чтения и записи
func readStdInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		return input
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return ""
}

func readStdInputSlice() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	var i int
	for scanner.Scan() {
		i++
		input = append(input, scanner.Text())
		numberOfPasswords, _ := strconv.Atoi(input[0])
		if numberOfPasswords*3 == i-1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return input
}

func writeStdOutput(output ...string) {
	for _, val := range output {
		fmt.Fprint(os.Stdout, val+"\n")
	}
}

func writeStdOutputSlice(output []string) {
	for _, val := range output {
		fmt.Fprint(os.Stdout, val+"\n")
	}
}
