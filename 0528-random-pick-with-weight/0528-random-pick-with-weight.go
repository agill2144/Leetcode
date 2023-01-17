type Solution struct {
    runningSums []int
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
        runningSums: rSums,
        total: total,
    }
}


func (this *Solution) PickIndex() int {
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := this.total
    target := rand.Intn(max - min + 1) + min

    left := 0
    right := len(this.runningSums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        if this.runningSums[mid] == target {return mid}
        if target < this.runningSums[mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return left
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */