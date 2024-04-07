/*
    approach: brute force
    - for each price being added
    - we need to tell how many CONSECUTIVE days, price is < current price
    - we need to go back as far as we can
    - until we run into a price > current day price
    - example; current day = 85
    - prev days = [100,65,75,82,81,76,70]
    - go back until ^
    - we stop at idx 0,
    - therefore number of days price was less than or equal to the price of curr day is
    - 6 + 1(current day) = 7

    time = o(n) for each Next() call
*/
type StockSpanner struct {
    prices []int    
}


func Constructor() StockSpanner {
    return StockSpanner{
        prices: []int{},
    }    
}


func (this *StockSpanner) Next(price int) int {
    this.prices = append(this.prices, price)
    lastIdx := len(this.prices)-1
    i := lastIdx
    for i >= 0 {
        if this.prices[i] > price {break}
        i--
    }
    return lastIdx-i
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */