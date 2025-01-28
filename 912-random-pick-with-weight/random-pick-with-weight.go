type Solution struct {
    compressedIdxs []int
}


func Constructor(w []int) Solution {
    rSum := 0
    compressedIdxs := []int{}
    for i := 0; i < len(w); i++ {
        rSum += w[i]
        compressedIdxs = append(compressedIdxs, rSum)
    }
    return Solution{compressedIdxs}
}


func (this *Solution) PickIndex() int {
    n := len(this.compressedIdxs)
    totalSum := this.compressedIdxs[n-1]
    left := 0
    right := n-1
    res := -1
    r := rand.Intn(totalSum)
    for left <= right {
        mid := left + (right-left)/2
        if this.compressedIdxs[mid] > r {
            res = mid
            right = mid-1
        }  else {
            left = mid+1
        }
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */