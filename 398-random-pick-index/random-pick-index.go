type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{nums}
}


func (this *Solution) Pick(target int) int {
    idx := -1
    count := 0
    for i := 0; i < len(this.nums); i++ {
        if this.nums[i] == target {
            count++
            if rand.IntN(count) == 0 {
                idx = i
            }
        }
    }
    return idx
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */