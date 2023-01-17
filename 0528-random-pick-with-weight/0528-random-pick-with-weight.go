type Solution struct {
    sums []int
}


func Constructor(w []int) Solution {
    total := 0
    rSums := make([]int, len(w))
    for i := 0; i < len(w); i++ {
        total += w[i]
        rSums[i] = total
    }
    return Solution{
        sums: rSums,
    }
}


func (this *Solution) PickIndex() int {
    n := len(this.sums)
    if n == 0 {return -1}
    
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := this.sums[n-1]
    target := rand.Intn(max - min + 1) + min

    left := 0
    right := n-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if this.sums[mid] == target || (target < this.sums[mid] && (mid == 0 || target > this.sums[mid-1])) {
            return mid
        } else if target > this.sums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */