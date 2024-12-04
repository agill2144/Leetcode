type Solution struct {
    compressedRanges []int
}


func Constructor(w []int) Solution {
    c := []int{w[0]}    
    for i := 1; i < len(w); i++ {
        c = append(c, w[i] + c[i-1])
    }
    return Solution{c}
}


func (this *Solution) PickIndex() int {
    n := len(this.compressedRanges)
    maxRange := this.compressedRanges[n-1]
    r := rand.Intn(maxRange)
    ans := -1
    left := 0
    right := n-1
    // mid position tells us the ending of range at mid idx
    // EXCLUDING mid value
    // if mid idx(2) val is 3
    // it means, the range at mid pos ends at 3( does not include 3)
    // we want to a smallest such mid where r IN the range (not equal to, but r < mid range )
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