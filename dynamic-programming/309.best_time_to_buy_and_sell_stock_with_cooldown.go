package dynamic_programming

import (
	"math"
)

func maxProfit309(prices []int) int {
	sold, rest := 0, 0
	hold := math.MinInt32
	for _, price := range prices {
		preSold := sold
		hold = max(hold, rest - price)
		sold = hold + price
		rest = max(rest, preSold)
	}
	return max(sold, rest)
}