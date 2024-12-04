
type Solution struct {
    // compressedIdxRanges represents the repetition of each index as ranges in a compact form.
    compressedIdxRanges []int
}


func Constructor(w []int) Solution {
    compressedIdxRanges := []int{}
    for i := 0; i < len(w); i++ {
        compressedIdxRanges = append(compressedIdxRanges,w[i])
        if i-1 >= 0 {compressedIdxRanges[i] += compressedIdxRanges[i-1]}
    }
    return Solution{compressedIdxRanges}
}


func (this *Solution) PickIndex() int {
    n := len(this.compressedIdxRanges)
    r := rand.Intn(this.compressedIdxRanges[n-1])
    left := 0
    right := n-1
    ans := -1
    // r is our target and we need to find 
    // the range where our target lies in
    // when is mid our potential ans?
    // r = 4
    // when mid value is 3 and mid+1 value is 7
    // meaning range from mid is from [3 -> 7)
    // from 3 to 7 excluding 7
    // therefore 4 falls under this range, this is a potential answer, save it and search left
    // because on the right side, the range will increase as the cumulative sum increases
    for left <= right {
        mid := left + (right-left)/2
        if r < this.compressedIdxRanges[mid]  {
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