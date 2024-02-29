/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction
(i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
Note that you cannot sell a stock before you buy one.
*/
// time: 2018-12-28

package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("prices of stcks ", prices)
	profit := maxProfit(prices)
	fmt.Println("Maximum Profit From stock ", profit)
	// profit1 := maxProfit1(prices)
	// fmt.Println("Maximum Profit From stock 1 ", profit1)
	// profit2 := maxProfit2(prices)
	// fmt.Println("Maximum Profit From stock 1 ", profit2)
}

// prices = [7,1,5,3,6,4]
//121 leetcode
// dynamic programming
// for every price, find the max profit, then record the current minimum price.
// time complexity: O(n)
// space complexity: O(1)
func maxProfit(prices []int) int {
	n := len(prices)
	if 0 == n || 1 == n {
		return 0
	}
	var (
		res      int
		minPrice = prices[0]
	)
	for _, price := range prices {
		if price-minPrice > res {
			res = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return res
}

// prices = [7,1,5,3,6,4]
//O(n) Time , O(1) Space (for two pointers solution)//sliding window
func maxProfit1(prices []int) int {
	left := 0
	bestProfit := 0
	for right := 1; right < len(prices); right++ {
		//fmt.Println("Iteration ", right)
		if prices[left] > prices[right] {
			//fmt.Println("at", left, "(", prices[left], ")", "GREATER THAN ", "at", right, "(", prices[right], ")")
			left = right
			//fmt.Println("left is greter than right ", left, right)
		} else {
			//fmt.Println("at", left, "(", prices[left], ")", "NOT GREATER THAN ", "at", right, "(", prices[right], ")")
			bestProfit = max(bestProfit, prices[right]-prices[left])
			//fmt.Println("Best profit updated ", bestProfit)
		}
	}
	return bestProfit
}

//update bestprofit/maxprofit if current profit is higher than the previous profit
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// brute force
// time complexity: O(n^2)
// space complexity: O(1)
func maxProfit2(prices []int) int {
	n := len(prices)
	if 0 == n || 1 == n {
		return 0
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > res {
				res = prices[i] - prices[j]
			}
		}
	}
	return res
}
