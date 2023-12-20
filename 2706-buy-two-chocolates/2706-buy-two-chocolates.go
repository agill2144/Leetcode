func buyChoco(prices []int, money int) int {
    smallest := math.MaxInt64
    second := math.MaxInt64
    for i := 0; i < len(prices); i++ {
        if prices[i] < smallest {
            second = smallest
            smallest = prices[i]
        } else if prices[i] < second {
            second = prices[i]
        }
    }
    totalCost := smallest + second
    if totalCost > money {return money}
    return money-totalCost
}