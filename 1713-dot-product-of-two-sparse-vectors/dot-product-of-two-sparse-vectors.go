type SparseVector struct {
    idx map[int]int // {idxFromNums: valueAtThatIdx}
}

func Constructor(nums []int) SparseVector {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        idx[i] = nums[i]
    }
    return SparseVector{idx}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    if len(vec.idx) < len(this.idx) {return vec.dotProduct(*this)}
    sum := 0
    for i, val := range this.idx {
        val2 := vec.idx[i]
        if val2 != 0 {
            sum += (val*val2)
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