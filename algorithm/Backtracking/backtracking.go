package backtracking

import (
	"fmt"
	"sort"
)

// 组合
func Combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var back func(n, k, startIndex int)
	back = func(n, k, startIndex int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := startIndex; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			back(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	back(n, k, 1)

	return res
}

// 组合总和 III 找出所有相加之和为 n 的 k 个数的组合
func CombinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int

	var back func(k, n, startIndex int)
	back = func(k, n, startIndex int) {
		if n < 0 {
			return
		}
		if len(path) == k {
			if n == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}

		for i := startIndex; i <= 9; i++ {
			path = append(path, i)
			back(k, n-i, i+1)
			path = path[:len(path)-1]
		}
	}

	back(k, n, 1)

	return res
}

// 电话号码的字母组合
func LetterCombinations(digits string) []string {
	var res []string
	if digits == "" {
		return res
	}
	mp := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	var back func(curStr string, i int)
	back = func(curStr string, i int) {
		if i == len(digits) {
			res = append(res, curStr)
			return
		}

		for _, s := range mp[digits[i]] {
			back(curStr+string(s), i+1)
		}
	}

	back("", 0)

	return res
}

// 组合总和
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}

	sort.Ints(candidates)
	var path []int

	var back func(index, sum int)

	back = func(index, sum int) {
		if sum < 0 {
			return
		}
		if sum == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			back(i, sum-candidates[i])
			path = path[:len(path)-1]
		}
	}

	back(0, target)

	return res
}

// 组合总和2
func CombinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}

	sort.Ints(candidates)
	var path []int

	var back func(startIndex, sum int)
	back = func(startIndex, t int) {
		if t < 0 {
			return
		}
		if t == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := startIndex; i < len(candidates); i++ {

			if i > startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			back(i+1, t-candidates[i])
			path = path[:len(path)-1]
		}
	}

	back(0, target)

	return res
}

// 切割回文串
func Partition(s string) [][]string {
	var res [][]string
	if s == "" {
		return res
	}

	var path []string
	var back func(startIndex int)

	back = func(startIndex int) {
		if startIndex == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}

		for i := startIndex; i < len(s); i++ {
			curStr := s[startIndex : i+1]
			fmt.Println(curStr)
			if isPalindrome(curStr) {
				path = append(path, curStr)
				back(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	back(0)

	return res
}

func isPalindrome(s string) bool {

	for l, r := 0, len(s)-1; l <= r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}

	return true
}
