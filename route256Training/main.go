package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
)

func main() {
	input := readStdInputSlice()
  // input := []string{"3", "3", "cba", "aaz", "2", "bc", "ab", "1", "a", "a"}
  passwordChecks := originalPassword(input)
  writeStdOutputSlice(passwordChecks)
}

// Задание 4. Тренировочный раунд.
func originalPassword(inputData []string) []string {
  var answer []string
  numberOfPasswords, _ := strconv.Atoi(inputData[0])
  for i:=0; i<numberOfPasswords*3; i=i+3 {
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
      if firstHalfLetter<secondHalfLetter {
        return true
      }
    }
  }
  return false
}

// Задание 3. Тренировочный раунд. Слишком медленно.
func happyOrder(orderNum string) string {
  orderNumInt, _ := strconv.Atoi(orderNum)
  for i:=orderNumInt ; ; i++ {
      if isPalindrom(i) {
        return strconv.Itoa(i-orderNumInt)
      }
  }
  return "No palindrom!"
}

func isPalindrom(num int) bool {
  numStr := strconv.Itoa(num)
  var reverseNumStr string
  for i:=len(numStr)-1; i>=0; i-- {
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


// Задание 2. Тренировочный раунд.
//TODO: Оптимизировать
func forgottenPincode(numbers string) []string {
  var pinOptions []string = strings.Split(numbers, "")
  var pincodeCombinations []string

  sortedPinOptionsInts := sortSlice(pinOptions)
  pincodeCombinations = append(pincodeCombinations, sliceItoa(sortedPinOptionsInts, 0,1,2))
  pincodeCombinations = append(pincodeCombinations, sliceItoa(sortedPinOptionsInts, 0,2,1))
  pincodeCombinations = append(pincodeCombinations, sliceItoa(sortedPinOptionsInts, 1,0,2))
  pincodeCombinations = append(pincodeCombinations, sliceItoa(sortedPinOptionsInts, 1,2,0))
  pincodeCombinations = append(pincodeCombinations, sliceItoa(sortedPinOptionsInts, 2,0,1))
  pincodeCombinations = append(pincodeCombinations, sliceItoa(sortedPinOptionsInts, 2,1,0))

  return pincodeCombinations
}

func sortSlice(stringPincode []string) []int {
  var ints []int
  for _, v := range stringPincode {
    num, _ := strconv.Atoi(v)
    ints = append(ints, num)
  }

  for i:=0; i<len(ints)-1; i++ {
    if ints[i] > ints[i+1] {
      ints[i], ints[i+1] = ints[i+1], ints[i]
    }
  }
  for i:=0; i<len(ints)-1; i++ {
    if ints[i] > ints[i+1] {
      ints[i], ints[i+1] = ints[i+1], ints[i]
    }
  }
  return ints
}

func sliceItoa(ints []int, num1, num2, num3 int) string {
  return strconv.Itoa(ints[num1]) + strconv.Itoa(ints[num2]) + strconv.Itoa(ints[num3])
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

func writeStdOutputSlice(output []string) {
  for _, val := range output {
      fmt.Fprint(os.Stdout, val + "\n")
  }
}

func writeStdOutputString(output string) {
  fmt.Fprint(os.Stdout, output + "\n")
}
