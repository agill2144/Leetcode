func maxProfit(prices []int) int {
    n := len(prices)
    buyPrices := make([]int, n)
    buyPrices[0] = prices[0]
    for i := 1; i < n; i++ {
        buyPrices[i] = min(buyPrices[i-1], prices[i-1])
    }
    profit := 0
    for i := 0; i < n; i++ {
        profit = max(profit, prices[i]-buyPrices[i])
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