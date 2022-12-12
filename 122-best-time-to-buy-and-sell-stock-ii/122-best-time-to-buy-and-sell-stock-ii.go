
/*
    We can do as many transactions as possible...
    Greedy: This is when we will compare buying on i and selling on i+1 price
    if there is any profit, add to total profit..
    
    Question becomes, if you sold it on a day, can you buy it again after selling it?
    The answer is not very clear but it is true, you can buy the stock after selling it
    
    Therefore we can apply greedy..
    buy and then sell
    and if selling created a profit, add to totalProfit
    then buy again on the same day and sell it again and repeat

    time: o(n)
    space: o(1)

*/
func maxProfit(prices []int) int {
    totalProfit := 0
    for i := 1; i < len(prices); i++ {
        buy := prices[i-1]
        sell := prices[i]
        profit := sell - buy 
        if profit > 0 {
            totalProfit += profit
        }
    }
    
    return totalProfit
}