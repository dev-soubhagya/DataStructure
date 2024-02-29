package main

import (
	"fmt"
	"math"
)

// To get understand
// https://www.codingninjas.com/codestudio/library/best-time-to-buy-and-sell-stock-iii

func main() {
	stockprices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	profit := maxProfit3(stockprices)
	fmt.Println("Profit got ", profit)
}
func maxProfit2(prices []int) int {
	t1_b, t1_s, t2_b, t2_s := math.MinInt32, 0, math.MinInt32, 0

	for _, v := range prices {
		t1_b = max(t1_b, 0-v)
		t1_s = max(t1_s, t1_b+v)
		t2_b = max(t2_b, t1_s-v)
		t2_s = max(t2_s, t2_b+v)
	}
	return t2_s
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return y
}

//giving wrong output
func maxProfit(prices []int) int {
	g, f := make([]int, len(prices)), make([]int, len(prices))
	ans, n, low := 0, len(prices), prices[0]

	for i := 1; i < n; i++ {
		low = min(low, prices[i])
		f[i] = max(f[i-1], prices[i]-low)
	}
	high := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		high = max(high, prices[i])
		g[i] = max(g[i+1], high-prices[i])
	}

	for i := 0; i < n; i++ {
		ans = max(ans, f[i]+g[i])
	}
	return ans
}

//Dynamic Programming Approach
func maxProfit3(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}

	l := make([]int, n)
	r := make([]int, n)
	l_min := prices[0]
	r_max := prices[n-1]
	tmp := 0
	for i := 1; i < n; i++ {
		if prices[i] < l_min {
			l_min = prices[i]
		}
		tmp = prices[i] - l_min
		if tmp > l[i-1] {
			l[i] = tmp
		} else {
			l[i] = l[i-1]
		}
	}

	for i := n - 2; i > -1; i-- {
		if prices[i] > r_max {
			r_max = prices[i]
		}
		tmp = r_max - prices[i]
		if tmp > r[i+1] {
			r[i] = tmp
		} else {
			r[i] = r[i+1]
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		tmp = l[i] + r[i]
		if tmp > result {
			result = tmp
		}
	}

	return result
}

//DP

func maxProfit4(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
	}

	// dp[k][i] = max(dp[k][i-1], prices[i] - prices[j] + dp[k-1][j-1])
	for k := 1; k < 3; k++ {
		min_price := prices[0]
		for i := 1; i < n+1; i++ {
			min_price = Min(min_price, prices[i-1]-dp[k-1][i-1])
			dp[k][i] = Max(dp[k][i-1], prices[i-1]-min_price)
		}
	}

	return dp[2][n]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//Simple Solution

func maxProfit5(prices []int) int {
	left := make([]int, len(prices))
	right := make([]int, len(prices))
	l := prices[0]
	r := prices[len(prices)-1]
	for i := 0; i < len(prices); i++ {
		min := Minn(prices[i], l)
		left[i] = prices[i] - min
		if i != 0 {
			left[i] = Maxx(left[i], left[i-1])
		}
		l = min
	}

	for i := len(prices) - 1; i >= 0; i-- {
		max := Maxx(prices[i], r)
		right[i] = max - prices[i]
		r = max
	}
	ans := 0
	for i := 0; i < len(left); i++ {
		tem := left[i]
		if i != len(prices)-1 {
			tem += right[i+1]
		}
		ans = Maxx(ans, tem)
	}
	return ans
}

func Maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Minn(a, b int) int {
	if a < b {
		return a
	}
	return b
}
