/*
    if you know the numbers are from 0 to n
    then we can calculate the sum of all numbers from 0 to n using a formula
    - (n(n+1))/2
    
    the above gives total EXPECTED sum
    now we can calc the acutal sum and return the diff
    - the diff is the missing number
*/
func missingNumber(nums []int) int {
    n := len(nums)
    expected := (n * (n+1))/2
    actual := 0
    for i := 0; i < len(nums); i++ {
        actual += nums[i]
    }
    return expected-actual
}