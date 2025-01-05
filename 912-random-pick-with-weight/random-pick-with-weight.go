type Solution struct {
    cmpIdxs []int
}


func Constructor(w []int) Solution {
    sum := 0
    cmpIdxs := []int{}    
    for i := 0; i < len(w); i++ {
        sum += w[i]
        cmpIdxs = append(cmpIdxs, sum)
    }
    return Solution{cmpIdxs}
}


func (this *Solution) PickIndex() int {
    n := len(this.cmpIdxs)
    end := this.cmpIdxs[n-1]
    r := rand.Intn(end)
    left := 0
    right := n-1
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.cmpIdxs[mid] <= r {
            left = mid+1
        } else {
            res = mid
            right = mid-1
        }
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */