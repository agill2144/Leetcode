func maxProfit(prices []int) int {
    n := len(prices)
    suffix := make([]int, n)
    for i := n-2; i >= 0; i-- {
        next := prices[i+1]
        maxSoFar := suffix[i+1]
        suffix[i] = max(maxSoFar, next)        
    }
    maxProfit := 0
    for i := 0; i < len(prices); i++ {
        buy := prices[i]
        sell := suffix[i]
        maxProfit = max(maxProfit, sell-buy)
    }
    return maxProfit
}