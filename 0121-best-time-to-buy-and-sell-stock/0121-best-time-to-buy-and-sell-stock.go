func maxProfit(prices []int) int {
    buyPrices := make([]int, len(prices))
    buyPrices[0] = prices[0]
    for i := 1; i < len(prices); i++ {
        minBuyPrice := prices[i]
        if buyPrices[i-1] < minBuyPrice {
            minBuyPrice = buyPrices[i-1]
        }
        buyPrices[i] = minBuyPrice
    }
    maxProfit := 0
    for i := 0; i < len(prices); i++ {
        sell := prices[i]
        buy := buyPrices[i]
        profit := sell-buy
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}