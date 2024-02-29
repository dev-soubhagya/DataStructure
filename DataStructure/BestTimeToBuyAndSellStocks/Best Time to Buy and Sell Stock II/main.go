package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

//122 leetcode
func main() {
	stockprices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(stockprices)
	// profit := maxProfit1(stockprices)
	// profit := maxProfit2(stockprices)
	//profit := maxProfit3(stockprices)
	fmt.Println("After Buy And sell stocks the maximum profit : ", profit)
}

//best solution
// O(n) Time, O(1) Space. solve by Greedy Approach
func maxProfit(prices []int) int {
	profit := 0
	// no profit available at index 0, so start at index 1
	for i := 1; i < len(prices); i++ {
		// Only when the current price is higher than the previous
		// can we make a profitable sale.
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
		// The else case is handled implicitly by incrementing i.
		// When prices[i] is lower than prices[i-1], we just skip
		// the sale.
		//
		// If this is confusing, try keeping a `min` variable yourself,
		// and you'll soon realise it's redundant.
	}
	return profit
}

// Peak Valley Approach
func maxProfit1(prices []int) int {

	prices = append(prices, -1) // append -1 is for final loop to add the last valley-peak profit
	profit := 0
	valley := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i-1] > prices[i] {

			profit += prices[i-1] - valley
			valley = prices[i]

		}
	}

	return profit

}

//Implementation by botoom-up DP + iteration:
func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func maxProfit2(prices []int) int {
	cur_hold, cur_not_hold := math.MinInt64, 0
	for _, stock_price := range prices {
		prev_hold, prev_not_hold := cur_hold, cur_not_hold
		// either keep hold, or buy in stock today at stock price
		cur_hold = Max(prev_hold, prev_not_hold-stock_price)
		// either keep not-hold, or sell out stock today at stock price
		cur_not_hold = Max(prev_not_hold, prev_hold+stock_price)
	}
	return cur_not_hold
}

// Implementation with Top down DP + recursion
func maxProfit3(prices []int) int {
	// inner function prototype
	var trade func(day_d int) (int, int)
	trade = func(day_d int) (int, int) {
		if day_d == 0 {
			return -prices[day_d], 0
		}
		prev_hold, prev_not_hold := trade(day_d - 1)
		hold := Max(prev_hold, prev_not_hold-prices[day_d])
		not_hold := Max(prev_not_hold, prev_hold+prices[day_d])
		return hold, not_hold
	}
	last_day := len(prices) - 1
	_, final_not_hold := trade(last_day)
	return final_not_hold
}
