func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    if k == 0 {return}
    reverse(0, n-1, nums)
    reverse(0, k-1, nums)
    reverse(k, n-1, nums)
}

func reverse(left, right int, nums []int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}

/*
    end = [5,6,7,1,2,3,4]

    [1,2,3,4,5,6,7]

    (0, n-1)
    [7,6,5,4,3,2,1]

    (0, k-1)
    [5,6,7,4,3,2,1]

    (k, n-1)
    [5,6,7,1,2,3,4]

*/