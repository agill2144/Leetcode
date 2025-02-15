type ProductOfNumbers struct {
    left int
    prods []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        left: 0,
        prods: []int{1},
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    last := 1
    if len(this.prods) > 0 {
        last = this.prods[len(this.prods)-1]
    }
    if num == 0 {
        this.prods = []int{1}
    } else {
        this.prods = append(this.prods, last*num)
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    n := len(this.prods)
    if k >= n {return 0}
    return this.prods[n-1] / this.prods[n-k-1]
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */