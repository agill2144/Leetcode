type SparseVector struct {
    items [][]int // [ [idx, val] ]
}

func Constructor(nums []int) SparseVector {
    items := [][]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {continue}
        items = append(items, []int{i, nums[i]})
    }
    return SparseVector{items}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    if len(vec.items) < len(this.items) {
        return vec.dotProduct(*this)
    }
    total := 0
    for i := 0; i < len(this.items); i++ {
        idx, val := this.items[i][0], this.items[i][1]
        val2 := binarySearch(vec.items, idx)
        if val2 != -1 {
            total += (val * val2)
        }
    }
    return total
}

func binarySearch(items [][]int, targetIdx int) int {
    left := 0
    right := len(items)-1
    for left <= right {
        mid := left + (right-left)/2
        midIdx, midVal := items[mid][0], items[mid][1]
        if midIdx == targetIdx {
            return midVal
        } else if targetIdx > midIdx {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */