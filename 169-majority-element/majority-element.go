// kadanes voting algo
func majorityElement(nums []int) int {
    candidate := math.MinInt64
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count = 1
            candidate = nums[i]
        } else if nums[i] == candidate {
            count++
        } else {
            count--
            if count < 0 {count = 0}
        }
    }
    // verfiy candidtate
    count = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == candidate {
            count++
        }
    }
    if count >= (len(nums)/2)+1 {return candidate}
    return -1
}