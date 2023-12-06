func maxProfit(prices []int) int {
    out := 0
    nextLargerNumber := prices[len(prices)-1]
    for i := len(prices)-2; i >= 0; i-- {
        profit := nextLargerNumber-prices[i]
        out = max(out, profit)
        nextLargerNumber = max(prices[i], nextLargerNumber)
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}