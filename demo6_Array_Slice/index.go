package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var numbers = make([]int, 0, 5) //type ,len,cap ขา ensurecapacity จะเพิ่ม cap ตามที่ระบึตอนแรก
	// var numbers2 []int
	// showSlice(numbers2)
	// numbers2 = append(numbers2, 1)
	// showSlice(numbers2)
	// subSlice()
	addStrings("123", "456")
}

// Nofix size pointer with Array len จำนวนสมาชิก, cap จุสูงสุดได้เท่าไหร่ (list)
func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func basicArr() {
	var arrInt [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 [10]string
	for _, v := range arrInt {
		fmt.Print(v, ",")
	}
	arr2[0], arr2[1], arr2[2] = "A", "B", "C"
	fmt.Print(arr2)
}

func subSlice() {
	var numbers = []int{1, 2, 3, 4, 5}
	showSlice(numbers)

	//leading remove ลบด้านหน้า ลบตั้งแต่ตัวที่3
	numbers = numbers[3:]
	showSlice(numbers)
	numbers = []int{1, 2, 3, 4, 5}
	//trailing remove ลบจากด้านหลัง
	numbers = numbers[0 : len(numbers)-2]
	showSlice(numbers)
	numbers = []int{1, 2, 3, 4, 5, 6, 6, 7}
	removeDuplicates(numbers)
	fmt.Println(numbers)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mySet := make(map[int]bool)
	var run int = 0
	for i, v := range nums {
		if !mySet[v] {
			mySet[v] = true
			nums[run] = nums[i]
			run++
		}
	}
	nums = nums[:run]
	return run
}

func removeByIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...) //เอาตัวแรกไปจนถึง index มาต่อกับหลัง index ถึงตัวท้าย ...to return as member
}

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0
	result := ""
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		result = strconv.Itoa(sum%10) + result
		carry = sum / 10
	}
	return result
}
