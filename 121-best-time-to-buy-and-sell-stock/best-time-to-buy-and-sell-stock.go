func maxProfit(prices []int) int {
    n := len(prices)
    greatestSellPrices := make([]int, n)
    for i := n-2; i >= 0; i-- {
        nextSellPrice := prices[i+1]
        greatestSoFar := greatestSellPrices[i+1]
        greatestSellPrices[i] = max(nextSellPrice, greatestSoFar)
    }
    maxEarn := 0
    for i := 0; i < n; i++ {
        maxEarn = max(maxEarn, greatestSellPrices[i]-prices[i])
    }
    return maxEarn
}