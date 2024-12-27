func maxProfit(prices []int) int {
    bestBuyPrices := make([]int, len(prices))
    bestBuyPrices[0] = prices[0]
    for i := 1; i < len(prices); i++ {
        bestBuyPrices[i] = min(prices[i], bestBuyPrices[i-1])
    }
    maxPr := 0
    for i := 0; i < len(prices); i++ {
        maxPr = max(maxPr, prices[i]-bestBuyPrices[i])
    }
    return maxPr
}