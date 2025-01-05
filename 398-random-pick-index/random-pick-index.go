type Solution struct {
    nums []int    
}


func Constructor(nums []int) Solution {
    return Solution{nums}    
}


func (this *Solution) Pick(target int) int {
    count := 0
    res := -1
    for i := 0; i < len(this.nums); i++ {
        if this.nums[i] == target {
            count++
            if rand.IntN(count) == 0 {
                res = i
            }
        }
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */