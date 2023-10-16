func maxProfit(prices []int) int {
    total := 0
    for i := 1 ; i < len(prices); i++ {
        buy := prices[i-1]
        sell := prices[i]
        if sell-buy > 0 {total += sell-buy }
    }
    return total
}