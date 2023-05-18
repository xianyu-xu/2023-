package dynamicProgramming

import (
	"fmt"
	"math"
)

//-------------------01 背包--------------------
/*
确定dp数组（dp table）以及下标的含义
确定递推公式
dp数组如何初始化
确定遍历顺序
举例推导dp数组
*/

// MinCostClimbingStairs 746. 使用最小花费爬楼梯
func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	if len(cost) < 2 {
		return 0
	}

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	fmt.Println(dp)
	return dp[len(cost)]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// UniquePaths 62. 不同路径
func UniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)

	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// UniquePathsWithObstacles 63. 不同路径 II
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	if m == 0 || n == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, m)

	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// IntegerBreak 343. 整数拆分
func IntegerBreak(n int) int {
	dp := make([]int, n+1)

	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// NumTrees 96. 不同的二叉搜索树
func NumTrees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

// BeiBaoErWei 背包问题二维数据
func BeiBaoErWei(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)

	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j > 0; j-- {
			if j >= weight[i] {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
			fmt.Println(dp, j)
		}
	}

	return dp[bagWeight]
}

// CanPartition 416. 分割等和子集
func CanPartition(nums []int) bool {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}

// LastStoneWeightII 1049. 最后一块石头的重量 II
func LastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2

	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[target] - dp[target]
}

// FindTargetSumWays 494. 目标和
func FindTargetSumWays(nums []int, target int) int {
	if (len(nums)+target)%2 != 0 {
		return 0
	}
	left := (len(nums) + target) / 2

	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	fmt.Println(dp)
	return dp[left]
}

// FindMaxForm 474. 一和零
func FindMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)

	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		x, y := 0, 0
		for _, str := range strs[i] {
			if str == '0' {
				x++
			} else {
				y++
			}
		}
		for j := m; j >= x; j-- {
			for k := n; k >= y; k-- {
				dp[j][k] = max(dp[j][k], dp[j-x][k-y]+1)
			}
		}
	}
	return dp[m][n]
}

//----------------------------完全背包----------------------

// Change 518. 零钱兑换 II
func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] += dp[i-coins[j]]
			}
		}
		fmt.Println(dp)
	}

	return dp[amount]
}

// CombinationSum4 377. 组合总和 Ⅳ
func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	fmt.Println(dp)
	return dp[target]

}

// ClimbStairs 爬楼梯
func ClimbStairs(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}

	fmt.Println(dp)

	return dp[n]
}

// CoinChange 零钱兑换
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	//for i := 0; i < len(coins); i++ {
	//	for j := coins[i]; j <= amount; j++ {
	//		if dp[j-coins[i]] != math.MaxInt {
	//			dp[j] = min(dp[j], dp[j-coins[i]]+1)
	//		}
	//	}
	//}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	fmt.Println(dp)

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// NumSquares 完全平方和
func NumSquares(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)

			fmt.Println(dp)
		}
	}

	return dp[n]
}
