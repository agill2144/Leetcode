func maxProfit(prices []int) int {
    minBuyPrice := math.MaxInt64
    profit := 0
    for i := 0; i < len(prices); i++ {
        sellPrice := prices[i]
        if sellPrice - minBuyPrice > profit {profit = sellPrice - minBuyPrice}
        if prices[i] < minBuyPrice {
            minBuyPrice = prices[i]
        }
    }
    return profit
}