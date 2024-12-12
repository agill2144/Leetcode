type SparseVector struct {
    nums [][]int // [ [val, idx] ]
}

func Constructor(nums []int) SparseVector {
    vals := [][]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {vals = append(vals, []int{nums[i], i})}
    }
    return SparseVector{vals}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    sum := 0
    v1, v2 := 0, 0

    for v1 < len(this.nums) && v2 < len(vec.nums) {
        idx1, idx2 := this.nums[v1][1], vec.nums[v2][1]

        if idx1 == idx2 {
            sum += this.nums[v1][0] * vec.nums[v2][0]
            v1++
            v2++
        } else if idx1 < idx2 {
            // Move v1 using binary search
            v1 = rightMostOnLeftSide(this.nums, v1, idx2)
            if this.nums[v1][1] != idx2 {v1++; v2++}
        } else {
            // Move v2 using binary search
            v2 = rightMostOnLeftSide(vec.nums, v2, idx1)
            if vec.nums[v2][1] != idx1 {v1++;v2++}
        }
    }

    return sum
}

func rightMostOnLeftSide(nums [][]int, left int, target int) int {
    right := len(nums) - 1
    ans := left

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid][1] <=  target {
            if nums[mid][1] == target {return mid}
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }

    return ans
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */