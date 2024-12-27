func maxProfit(prices []int) int {
    bestBuyPrice := prices[0]
    maxPr := 0
    for i := 1; i < len(prices); i++ {
        maxPr = max(maxPr, prices[i]-bestBuyPrice)
        bestBuyPrice = min(prices[i], bestBuyPrice)
    }
    return maxPr
}