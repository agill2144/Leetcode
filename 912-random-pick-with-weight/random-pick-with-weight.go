type Solution struct {
    compressedRanges []int
}


func Constructor(w []int) Solution {
    compressedRanges := make([]int, len(w))
    for i := 0; i < len(w); i++ {
        compressedRanges[i] = w[i]
        if i-1 >= 0 {compressedRanges[i] += compressedRanges[i-1]}
    }
    return Solution{compressedRanges}
}

// remember that values in compressed range indicate end idx of a range
// if compressedRanges[1] =  4
// it means, it covers until idx 4 from larger/expanded array
// but excluding idx 4
// therefore if r is < 4 (ex: 3), then yes this is potential ans, save it
// compressedRanges[1] = 2 ; ( meaning i cover up to idx 2 but not inculding 2)
// and if r = 3 or if r = 2; we are not in the range, therefore look right
func (this *Solution) PickIndex() int {
    n := len(this.compressedRanges)
    total := this.compressedRanges[n-1]
    r := rand.Intn(total)
    left := 0; right := n-1; ans := -1
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