type Solution struct {
    idxs map[int][]int
}


func Constructor(nums []int) Solution {
    idxs := map[int][]int{}    
    for i := 0; i < len(nums); i++ {
        idxs[nums[i]] = append(idxs[nums[i]], i)
    }
    return Solution{idxs}
}


func (this *Solution) Pick(target int) int {
    idxs := this.idxs[target]    
    r := rand.Intn(len(idxs))
    return idxs[r]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */