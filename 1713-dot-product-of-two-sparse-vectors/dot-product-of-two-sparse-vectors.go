type SparseVector struct {
    n int
    idxToVal map[int]int
}

func Constructor(nums []int) SparseVector {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {continue}
        m[i] = nums[i]
    }
    return SparseVector{len(nums), m}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    res := 0
    v1, v2 := 0, 0
    for v1 < this.n && v2 < vec.n {
        v1Val, ok1 := this.idxToVal[v1]
        v2Val, ok2 := vec.idxToVal[v2]
        if ok1 && ok2 {
            res += (v1Val * v2Val)
        }
        v1++; v2++
    }
    return res
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */