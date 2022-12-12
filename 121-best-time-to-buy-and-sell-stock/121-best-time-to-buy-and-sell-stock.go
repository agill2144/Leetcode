
/*
    since we have to return the largest possible profit,
    We want to buy it on the cheapest day possible and sell it on the day that gives us most profit
    
    We will start by implying that cheapest day is at idx 0
    so we will buy at price[0]
    
    Then loop over prices from 1, and check the profit if we were to sell on for ith price.
    (price[i]-buy) -- if this profit > max - then update max to the new number.
    
    However, because we implied that at 0th idx, the price is the cheapest, we need to also keep on checking
    if the current ith price is cheaper/smaller than the price we originally bought it for.
    If it is, save this as our new buy price and compare with rest of the selling prices down the line.
    
    Return the max at the end
    
    time: o(n)
    space: o(1)

*/
func maxProfit(prices []int) int {
    max := 0
    buy := prices[0]
    
    for i := 1; i < len(prices); i++ {
        sell := prices[i]
        profit := sell-buy
        if profit > max {
            max = profit
        }
        if sell < buy {
            buy = sell
        }   
    }
    
    return max
}
// Potential follow up - also print the buy and selling day that gets us the most profit
// func maxProfit(prices []int) int {
//     buyIdx := 0
//     buyDay := -1
//     sellDay := -1
//     max := 0
//     for i := 1; i < len(prices); i++ {
//         buy := prices[buyIdx]
//         sell := prices[i]
        
//         profit := sell - buy
//         if profit > 0 {
//             if profit > max {
//                 max = profit
//                 buyDay = buyIdx+1
//                 sellDay = i+1
//             }
//         }
//         if sell < prices[buyIdx] {
//             buyIdx = i
//         }
//     }
//     fmt.Println("Max profit would be made if we bought it on day", buyDay," and sold it on day ", sellDay)
//     return max
    
// }