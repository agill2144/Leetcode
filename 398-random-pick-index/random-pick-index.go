type Solution struct {
    idx map[int][]int // {val: [idx1, idx2, idx3]}
}


func Constructor(nums []int) Solution {
    m := map[int][]int{}
    for i := 0; i < len(nums); i++ {
        m[nums[i]] = append(m[nums[i]], i)
    }
    return Solution{idx:m}
}


func (this *Solution) Pick(target int) int {
    indices := this.idx[target]
    r := rand.Intn(len(indices))
    return indices[r]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */