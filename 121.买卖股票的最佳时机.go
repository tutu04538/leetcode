/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	n := len(prices)
	maxPrice, res := prices[n-1], 0
	for i := n - 2; i >= 0; i-- {
		res = max(res, maxPrice-prices[i])
		maxPrice = max(maxPrice, prices[i])
	}
	return res
}

// @lc code=end
