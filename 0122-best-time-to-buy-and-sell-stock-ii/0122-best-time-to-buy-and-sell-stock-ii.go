func maxProfit(prices []int) int {
    if len(prices) <= 1 {return 0}
    profit := 0
    for i := 1; i < len(prices); i++ {
        sell := prices[i]
        buy := prices[i-1]
        if sell-buy > 0 {profit+= sell-buy}
    }
    return profit
}