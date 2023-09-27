/*
    asking for something optimal, having multiple options
    
    approach: DP on stocks
    - we can only sell the stock at a given day if we have bought it before that given day
    - suppose we want to sell the stock on ith day, then it means we must have bought it before the ith day ( day 0 to day i-1 )
    - it makes sense to buy it on the cheapest day before selling on day i
        - therefore we need to find the cheapest price that we have seen before day i
        - this cheapest price acts as our "buying price"
    - for each ith selling day, we can search for a min number from 0 to i-1
        - this would be n^2 time
    - or we can keep a rolling minBuyPrice seen so far.
        - that is, instead of going and checking each and everytime
        - we can keep the min seen so far until i in a var "minBuyPrice"
        - we can calc profit ( sellPrice - minBuyPrice ) and save if this profit is better than a max profit maintained so far
        - since we are keeping a rolling min, we will update the minBuyPrice if current price is cheaper than earlier minBuyPrice
    
    time = o(n)
    space = o(1)
*/
func maxProfit(prices []int) int {
    if len(prices) == 0 {return 0}
    minBuyPrice := prices[0]
    profit := 0
    for i := 1; i < len(prices); i++ {
        currentProfit := prices[i]-minBuyPrice
        if currentProfit > profit {
            profit = currentProfit
        }
        // update the rolling min seen so far
        if prices[i] < minBuyPrice {
            minBuyPrice = prices[i]
        }
    }
    return profit
}