type Solution struct {
    prefix []int
}


func Constructor(w []int) Solution {
    prefix := make([]int, len(w))
    for i := 0; i < len(w); i++ {
        prefix[i] = w[i]
        if i-1 >= 0 {prefix[i] += prefix[i-1]}
    }
    return Solution{prefix}
}



func (this *Solution) PickIndex() int {
    n := len(this.prefix)
    r := rand.Intn(this.prefix[n-1])
    left := 0
    right := n-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        if this.prefix[mid] <= r {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans

}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */