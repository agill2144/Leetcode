func maxProfit(prices []int) int {
    buyPrice := prices[0]
    maxProfit := 0
    for i := 0; i < len(prices); i++ {
        sell := prices[i]
        profit := sell-buyPrice
        if profit > maxProfit {
            maxProfit = profit
        }
        if sell < buyPrice {
            buyPrice = sell
        }
    }
    return maxProfit
}