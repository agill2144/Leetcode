func maxProfit(prices []int) int {
    nextLargerNumber := make([]int, len(prices))
    for i := len(prices)-2; i >= 0; i-- {
        currNext := nextLargerNumber[i+1]
        nextNum := prices[i+1]
        nextLargerNumber[i] = max(currNext, nextNum)
    }
    
    out := 0
    for i := 0; i < len(prices); i++ {
        profit := nextLargerNumber[i]-prices[i]
        out = max(out, profit)
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}