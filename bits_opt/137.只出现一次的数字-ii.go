package main

import "fmt"

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	result := int32(0)
	for i := 0; i <= 31; i++ {
		bitCount := 0
		for j := 0; j < len(nums); j++ {
			if (nums[j]>>i)&1 == 1 {
				bitCount++
			}
		}
		if bitCount%3 != 0 {
			result |= 1 << i
		}
	}

	return int(result)
}

// @lc code=end

func main() {
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
