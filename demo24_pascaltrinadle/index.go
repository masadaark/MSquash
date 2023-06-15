package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1}
	fmt.Println(threeSum(nums))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func generate(numRows int) [][]int {
	newSlice := make([][]int, 0, numRows)
	newSlice = append(newSlice, []int{1})
	for i := 2; i <= numRows; i++ {
		i2 := i - 2
		inner := make([]int, 0, i)
		inner = append(inner, 1)
		for j := 0; j < i2; j++ {
			idx := j + 1
			sliceBefore := newSlice[i2]
			inner = append(inner, sliceBefore[idx]+sliceBefore[idx-1])
		}
		inner = append(inner, 1)
		newSlice = append(newSlice, inner)
	}
	return newSlice
}

func missingNumber(nums []int) int {
	n := len(nums) + 1
	summary := n * (n - 1) / 2
	for i := 0; i < n-1; i++ {
		summary -= nums[i]
	}
	return summary
}

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return len(nums)
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	return result
}
