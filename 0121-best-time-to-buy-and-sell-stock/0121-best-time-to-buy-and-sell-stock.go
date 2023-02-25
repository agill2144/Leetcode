func maxProfit(prices []int) int {
    max := 0
    buy := prices[0]
    for i := 1; i < len(prices); i++ {
        sell := prices[i]
        if sell-buy > max {max = sell-buy}
        if sell < buy {
            buy = sell
        }
    }
    return max
}