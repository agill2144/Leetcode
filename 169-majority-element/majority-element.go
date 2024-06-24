/*
    The majority element is the element that appears more than ⌊n / 2⌋ times
    therefore an element that has a count of > n/2, is the majority element
    
    Another famous majority counter algorithm that takes o(1) space is called "Mores voting algorithm"
    
    approach: mores voting using balancing-counter
    - the idea becomes if there is such majority element
    - then it must appear more than n/2 times
    - and we can use a counter to keep check whether such list contains a element that repeats > n/2 times
    - use counter to check if an array is balanced or not
    - start by taking the initial number as a "potential ans", therefore track
        - count = 1
        - ele = nums[0]
    - now loop over the array 
    - if nums[i] == ele ; count++
    - otherwise count--
    - when the array is kind of balanced
        - [1,1,1,2,3,4] , at the end count will become 0
        - 2 cancels 1, 3 cancels another 1, 4 cancels another 1
        - or
        - [1,2,3,4,5,6] , at the end count will become 0
        - each number cancels previous number
    - then the count will become 0
    - if there is FOR SURE a number that repeats more than n/2 times, it will leave the count as 1
        - if exactly more than half of the array is the same number,
        - then other numbers wont be able to cancel count to 0, example;
        - [1,2,1,3,4,1,1]
        c  1 0 1 0 1 0 1
        e  1 - 1 - 4 - 1
        c = count, e = element
        notice how 1 repeats more than n/2 times 
        n = 7; 7/2 = 3, so we need a number that repeats > 3 times
        1 repeats 4 times, therefore it kept the counter up
    - at anypoint when the count becomes 0, it means from idx0 till this idx,
    - there wasnt a majority element, if there was such majority up until this idx, then the count would have been > 0
    
    
    time = o(n)
    space = o(1)
    
    Follow up: How can we prove that there is exactly 1 element repeating more than n/2 times ?
    - current lc question gurantees 1 such element in all test cases, but what if that was not the case
    - eg; [1,2,3,4,5,5]
        c  1 0 1 0 1 2 
        e  1 - 1 - 5 5
    - 5 is not the majority element, how can we ensure whether the test case is not guranteed to container 1 majority element
    - well if the count at the is 0, then for sure nothing to check
    - otherwise, we need run another loop and check how many times 5 has repeated 
        - that would be 2 times
        - therefore its not a majority element
    
*/
func majorityElement(nums []int) int {
    ele := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count++
            ele = nums[i]
        } else if nums[i] == ele {
            count++
        } else {
            count--
        }
    }

    // if its not guranteed that there will always be such element that
    // appears more than n/2 times, check our final candidate $ele
    count = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele {
            count++
        }
    }
    if count > len(nums)/2 {return ele}
    return -1
}

/*
    majority element is a number that appears more than n/2 times
    approach: freq map
    - maintain a freq for each number
    - if at any point a number freq > n/2 times, we have found the majority element
    
    time = o(n)
    space = o(n)
*/
// func majorityElement(nums []int) int {
//     freqMap := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         freqMap[nums[i]]++
//         if val := freqMap[nums[i]]; val > len(nums)/2 {return nums[i]}
//     }
//     return -1
// }