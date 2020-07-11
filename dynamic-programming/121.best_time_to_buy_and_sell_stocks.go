package dynamic_programming

func maxProfit(prices []int) int {
	profit, sum := 0, 0
	for i := 1; i < len(prices); i++ {
		sum += prices[i] - prices[i-1]
		if sum < 0 {
			sum = 0
		}
		profit = max(profit, sum)
	}
	return profit
}
