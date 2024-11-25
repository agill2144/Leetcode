type SparseVector struct {
    idxToVal map[int]int
}

func Constructor(nums []int) SparseVector {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {continue}
        m[i] = nums[i]
    }
    return SparseVector{m}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    res := 0
    for idx, val := range this.idxToVal {
        vecVal, ok := vec.idxToVal[idx]
        if ok {
            res += (val*vecVal)
        }
    }
    return res
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */