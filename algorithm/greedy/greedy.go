package greedy

import (
	"math"
	"sort"
)

// FindContentChildren 455. 分发饼干
func FindContentChildren(g []int, s []int) int {
	res := 0
	if len(g) == 0 || len(s) == 0 {
		return res
	}

	index := len(s)-1

	sort.Ints(g)
	sort.Ints(s)

	for i := len(g)-1; i >= 0; i-- {
		for index >= 0 && s[index] >= g[i] {
			index--
			res++
			break
		}
	}

	return res
}

// WiggleMaxLength 376. 摆动序列
func WiggleMaxLength(nums []int) int {
	res := 1

	preDiff := 0
	curDiff := 0

	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1]-nums[i]
		if (preDiff >= 0 && curDiff < 0) || (preDiff <= 0 &&curDiff > 0) {
			res++
			preDiff = curDiff
		}
	}

	return res
}

// MaxSubArray 53. 最大子数组和
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := math.MinInt
	tmp := 0

	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if tmp > res {
			res = tmp
		}

		if tmp <= 0 {
			tmp = 0
		}
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxProfit 122. 买卖股票的最佳时机 II
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i]-prices[i-1]
		}
	}

	return res
}

// CanJump 55. 跳跃游戏
func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	jump := 0
	for i := 0; i <= jump; i++ {
		jump = max(i+nums[i], jump)
		if jump >= len(nums)-1 {
			return true
		}
	}

	return false
}















