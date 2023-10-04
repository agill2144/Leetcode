/*
    when selling on ith day, we only care about
    cheapest price seen so far from 0 to i-1
    - therefore we can keep track of cheapest price
    - in a rolling minPrice var
    
    time = o(n)
    space = o(1)
*/
func maxProfit(prices []int) int {
    if prices == nil {return 0}
    minPrice := prices[0]
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        profit := prices[i]-minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
        if prices[i] < minPrice {
            minPrice = prices[i]
        }
    }
    return maxProfit
}
/*
    we must buy the stock before selling on ith day
    - so if we want to maximize profit on the ith day
    - it makes sense to buy it on the cheapest possible day prior to ith day
    - so for each ith selling price, we need to know the cheapest price before ith day
    - i.e cheapest price on the left of ith position
    - we can create a prefix array that maintains the cheapest price on the left of i
    - and then we can use this prefix array to calc profit on ith day
    - which will be prices[i]-prefix[i]
        - prefix[i] will answer whats the cheapest price before ith position
    
    time o(n)
    space o(n)
*/
// func maxProfit(prices []int) int {
//     minBuyPrices := make([]int, len(prices))
//     minBuyPrices[0] = prices[0]
//     for i := 1; i < len(prices); i++ {
//         prevPrice := prices[i-1]
//         prevMin := minBuyPrices[i-1]
//         if prevPrice < prevMin {
//             minBuyPrices[i] = prevPrice
//         } else {
//             minBuyPrices[i] = prevMin
//         }
//     }
//     maxProfit := 0
//     for i := 1; i < len(prices); i++ {
//         buyPrice := minBuyPrices[i]
//         sellPrice := prices[i]
//         profit := sellPrice-buyPrice
//         if profit > maxProfit  {
//             maxProfit = profit
//         }
//     }
//     return maxProfit
    
// }