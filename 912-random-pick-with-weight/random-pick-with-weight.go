type Solution struct {
    compr []int
}


func Constructor(w []int) Solution {
    rSum := 0
    compr := []int{}
    for i := 0; i < len(w); i++ {
        rSum += w[i]
        compr = append(compr, rSum)
    }
    return Solution{compr}
}


func (this *Solution) PickIndex() int {
    n := len(this.compr)
    target := rand.Intn(this.compr[n-1])
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        start := 0
        end := this.compr[mid]
        if mid-1 >= 0 {start = this.compr[mid-1]}
        if target >= start && target < end {
            return mid
        }
        if target < start {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */