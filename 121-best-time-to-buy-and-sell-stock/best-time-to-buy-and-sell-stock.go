/*
    inorder to maximize earning
    we need the cheapest buy price
    and highest selling price
    ( but buy should happen before selling )
    that is, buy idx should be < sell idx

    - we can evalute each day as selling day
    - while keeping a rolling min buy value
    - start with a prices[0] as buy price
    - then loop from 1 to n-1
    - update max profit if there was a profit as needed
    - then update rolling min buy price if current day has a cheaper buy price
*/
func maxProfit(prices []int) int {
    if len(prices) == 1 {return 0}
    buy := prices[0]
    maxEarn := 0
    for i := 1; i < len(prices); i++ {
        maxEarn = max(maxEarn, prices[i]-buy)
        buy = min(buy, prices[i])
    }
    return maxEarn
}
// pre-compute the max selling prices from right side of the array ( like prefix array )
// then another pass to determine if buying on ith day, whats the max profit that can be earned
// for each ith buying price, we already have max selling price
// func maxProfit(prices []int) int {
//     n := len(prices)
//     greatestSellPrices := make([]int, n)
//     for i := n-2; i >= 0; i-- {
//         nextSellPrice := prices[i+1]
//         greatestSoFar := greatestSellPrices[i+1]
//         greatestSellPrices[i] = max(nextSellPrice, greatestSoFar)
//     }
//     maxEarn := 0
//     for i := 0; i < n; i++ {
//         maxEarn = max(maxEarn, greatestSellPrices[i]-prices[i])
//     }
//     return maxEarn
// }