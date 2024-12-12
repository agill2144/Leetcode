type SparseVector struct {
    idxToVal map[int]int
}

func Constructor(nums []int) SparseVector {
    idxToVal := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {idxToVal[i] = nums[i]}
    }
    return SparseVector{idxToVal}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    if len(vec.idxToVal) < len(this.idxToVal) {
        return vec.dotProduct(*this)
    } 
    sum := 0
    for idx, num := range this.idxToVal {
        num2, ok := vec.idxToVal[idx]
        if ok {
            sum += (num*num2)
        }
    }
    return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */