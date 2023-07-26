package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func IsPrime(n int) bool {
	//your code here
	if n == 2 || n == 1 {
		return true
	}
	for i := 1; i <= n; i++ {
		if i == 2 && n%2 == 0 {
			return false
		}

	}

	return true
}

func Solution(number int) string {
	// convert the number to a roman numeral
	result := ""
	sys := []string{" M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := strconv.Itoa(number)
	spl := strings.Split(strs, "")
	strs = spl[0]
	for i := 1; i < len(spl); i++ {
		strs += "0"
	}

	for i := 0; i < len(nums); i++ {
		if nm, _ := strconv.Atoi(strs); nm == nums[i] {
			result += sys[i]
			break
		}
	}
	spl2 := strings.Split(strs, "")
	for i := 0; i < len(spl2); i++ {
		st := strconv.Itoa(nums[i])
		if spls := strings.Split(st, ""); len(spl2) == len(spls) {
			for k := 1; k <= len(spls); k++ {
				result += sys[i]
			}
			break
		}
	}
	return result
}

func Persistence(n int) int {
	// your code
	count := 0
	its := strconv.Itoa(n)
	if len(its) == 1 {

		return 0
	}
	spl := strings.Split(its, "")
	result := ""
	number, _ := strconv.Atoi(spl[0])
	for j := 1; j <= len(spl)-1; j++ {
		convert, _ := strconv.Atoi(spl[j])
		number *= convert
	}
	count++
	convs := strconv.Itoa(number)
	result = convs
	for i := 1; len(result) != 1; i++ {
		count++
		spls := strings.Split(result, "")
		nums, _ := strconv.Atoi(spls[0])
		for j := 1; j <= len(spls)-1; j++ {
			convert, _ := strconv.Atoi(spls[j])
			nums *= convert
		}
		conv1 := strconv.Itoa(nums)
		result = conv1
	}
	return count

}
func FindEvenIndex(arr []int) int {

	vLeft := 0
	vRight := 0
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			vLeft = 0
		} else {
			for j := 0; j <= i-1; j++ {
				vLeft += arr[j]
			}
		}
		for k := i + 1; k <= len(arr)-1; k++ {
			vRight += arr[k]
		}
		if vLeft == vRight {
			return i
		}
		vLeft = 0
		vRight = 0
	}
	return -1
}
func toWeirdCase(str string) string {
	// Your code here and happy coding!
	weirdString := ""
	splW := strings.Split(str, " ")
	var spl2 [][]string
	for i := 0; i < len(splW); i++ {
		spl2 = append(spl2, []string{splW[i]})
	}

	for i := 0; i < len(spl2); i++ {
		spls := strings.Split(spl2[i][0], "")
		for j := 0; j <= len(spls)-1; j++ {

			if (j+1)%2 == 1 {
				weirdString += strings.ToUpper(spls[j])
			} else {
				weirdString += strings.ToLower(spls[j])

			}
		}
		if i != len(spl2)-1 {
			weirdString += " "

		}
	}

	return weirdString

}

func LongestConsec(strarr []string, k int) string {
	// your code
	cString := ""
	if len(strarr) == 0 || k > len(strarr) || k <= 0 {
		return ""
	}
	var csArr []string
	for i := 0; i <= len(strarr)-1; i++ {
		c := 0
		for j := i; j <= len(strarr)-1; j++ {
			c++
			cString += strarr[j]
			if c == k {
				break
			}
		}
		csArr = append(csArr, cString)
		cString = ""
	}

	cString = csArr[0]

	for i := 0; i < len(csArr); i++ {
		if len(csArr[i]) > len(cString) {
			cString = csArr[i]
		}
	}
	return cString
}

func duplicate_count(s1 string) int {
	count := 0
	splited := strings.Split(s1, "")
	newStr := []string{}
	for k := 0; k < len(splited); k++ {
		newStr = append(newStr, strings.ToLower(splited[k]))
	}
	sort.Strings(newStr)
	char := ""
	fmt.Println(newStr)
	for j := 0; j < len(newStr)-1; j++ {
		if j == 0 {
			if newStr[j] == newStr[j+1] {
				char = newStr[j]
				count++
				j++
			}
		} else {
			if newStr[j] == newStr[j+1] && newStr[j] != char {
				char = newStr[j]
				count++
				j++
			}
		}

	}

	return count //Instead of returning '1', you have to return integer 'i' as answer of solution.
}

func main() {
	fmt.Println(IsPrime(19))
	fmt.Println(Solution(1000))
	fmt.Println(Persistence(4))
	fmt.Println(toWeirdCase("abc"))
	duplicate_count("indivisibility")
	fmt.Println(LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2))
	fmt.Println(FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1}))
}
