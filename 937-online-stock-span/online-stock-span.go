type stackNode struct {
    idx int
    val int
}
type StockSpanner struct {
    idx int
    st []stackNode
}


func Constructor() StockSpanner {
    return StockSpanner{
        idx: 0,
        st: []stackNode{},
    }
}


func (this *StockSpanner) Next(price int) int {
    // ngl
    for len(this.st) != 0 && this.st[len(this.st)-1].val <= price {
        this.st = this.st[:len(this.st)-1]
    }
    ans := 0
    if len(this.st) != 0 {
        ans = this.idx - this.st[len(this.st)-1].idx        
    } else {
        
        ans = this.idx+1 
    } 
    this.st = append(this.st, stackNode{this.idx,price})
    this.idx++
    return ans
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */