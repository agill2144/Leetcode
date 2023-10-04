func maxProfit(prices []int) int {
    n := len(prices)
    profit := 0
    minBuyPrice := prices[0]
    for i := 1; i < n; i++ {
        profit = max(profit, prices[i]-minBuyPrice)
        minBuyPrice = min(minBuyPrice, prices[i])
    }
    return profit
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {return x}
    return y
}