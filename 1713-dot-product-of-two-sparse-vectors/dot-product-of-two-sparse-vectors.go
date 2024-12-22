type SparseVector struct {
    idxs map[int]int
}

func Constructor(nums []int) SparseVector {
    idxs := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {continue}
        idxs[i] = nums[i]
    }
    return SparseVector{idxs}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    if len(vec.idxs) < len(this.idxs) {
        return vec.dotProduct(*this)
    }
    total := 0
    for key, val := range this.idxs {
        val2, ok := vec.idxs[key]
        if ok {
            total += (val * val2)
        }
    }
    return total
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */