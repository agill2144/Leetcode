func missingElement(nums []int, k int) int {
    offSet := nums[0]
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        correctVal := mid+offSet
        missing := nums[mid]-correctVal
        if missing < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    missing := nums[right]-(right+offSet)
    // if at right ptr, we are missing x values
    // and lets assume x is a +ve value and is < k
    // this means on left of rightVal, there are x values missing
    // which means, the missing value is going to be k-x steps ahead of rightVal
    // for example; rightVal = 10; missing on left = 3; k = 5
    // if there are 3 values missing on left of 10, its obvious that we need to simply go 2 steps ahead of 10
    // why ? we need 5th missing, 3 are missing on left of 10, that leaves us with 2 ( 5-3 ; i.e k-missing )
    // then missing kth value is 10+2 = 12; rightVal + (k-missing) 
    // thats the intuition here
    // if there are x missing on left of right, how many more steps ahead of right val is our missing value ?
    // however many ever is remaining after taking out x from k
    return nums[right] + (k-missing)
}