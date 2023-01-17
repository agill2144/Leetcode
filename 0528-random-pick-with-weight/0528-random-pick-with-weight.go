type Solution struct {
    sums []int
    total int
}


func Constructor(w []int) Solution {
    total := 0
    rSums := []int{}
    for i := 0; i < len(w); i++ {
        total += w[i]
        rSums = append(rSums, total)
    }
    return Solution{
        sums: rSums,
        total: total,
    }
}


func (this *Solution) PickIndex() int {
    
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := this.total
    target := rand.Intn(max - min + 1) + min

    left := 0
    right := len(this.sums)-1
    
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