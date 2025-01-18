func maxProfit(prices []int) int {
    buy := prices[0]
    pr := 0
    for i := 0; i < len(prices); i++ {
        pr = max(pr, prices[i]-buy)
        buy = min(buy, prices[i])
    }
    return pr
}