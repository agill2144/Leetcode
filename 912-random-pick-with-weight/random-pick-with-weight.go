type Solution struct {
    compressedRanges []int
}


func Constructor(w []int) Solution {
    compressedRanges := make([]int, len(w))    
    for i := 0; i < len(w); i++ {
        compressedRanges[i] += w[i]
        if i-1 >= 0 {compressedRanges[i] += compressedRanges[i-1]}
    }
    return Solution{compressedRanges}
}


func (this *Solution) PickIndex() int {
    n := len(this.compressedRanges)
    r := rand.Intn(this.compressedRanges[n-1])
    ans := -1
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if r < this.compressedRanges[mid] {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */