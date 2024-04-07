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


    approach: monotonic stack
    - we need to look back at history ( stack hint )
    - if we had nested iterations, our inner loop start pos
    - depends on outter i's position ( ngr , ngl )
    - we need to go back until we run into a price > current price
        - next greater on left ( stack problem hint )
    - therefore we will apply classic NGL stack template
    - we need to find the number of days between current day and the day when ngl was pushed
    - therefore, whenever we push an item to stack, we will push a current day ( idx ) to it
    - we will track the current day as idx as a global var
    - when we need to find ngl
        - thats easy, remove all elements that are smaller than current price from top of stack
    - when we reach a ngl element
        - it will have idx ( i.e day ) as part of it
        - then finding the number of days should be easy;
        - newIdx - ngl idx
    - wont popping elements from stack mess up their idx ?
        - no, because those idx are constant, they represent day
        - meaning, price was. $x on day $idx
        - so even if we delete items from stack, the ngl will show correct day it was pushed on to stack
    - this way, each Next() call may not have to go thru ALL n days, it will be less on average compared to brute force

    time = o(n)
*/
type stackNode struct {
    val int
    idx int
}

type StockSpanner struct {
    st []stackNode
    idx int
}


func Constructor() StockSpanner {
    return StockSpanner{
        st: []stackNode{},
    }    
}


func (this *StockSpanner) Next(price int) int {
    // clean up / remove prices <= current price
    // we want ngl ( next greater on left )
    // therefore currPrice will remove anyone/all that are <= currPrice
    for len(this.st) != 0 && this.st[len(this.st)-1].val <= price {
        this.st = this.st[:len(this.st)-1]
    }

    days := 0
    if len(this.st) != 0 {
        // top is ngl, we want to use its idx/day
        days = this.idx - this.st[len(this.st)-1].idx
    } else {
        // if stack is empty, 
        // and there were 3 elements [1,2,3]
        // thats 3 days + current new day = 4
        days = this.idx+1
    }
    // push current pice
    this.st = append(this.st, stackNode{price, this.idx})
    this.idx++
    return days
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */