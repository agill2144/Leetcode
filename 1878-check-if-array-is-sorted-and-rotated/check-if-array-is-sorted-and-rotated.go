/*
    do we find the first min? and start traversing to check if rest is sorted from there?
    - no
    - ex: [1,2,2,3,4,5,1,1]
    - technically, idx 0 is a min
    - so is idx n-2
    - so which is the right min?
    - the correct min is where ith element > i+1
    - idx n-3 (5) and idx n-2 (1)
    - thats where the rotation point starts from...
    - therefore its not always the first min idx, we need to find the first rotation point
    - the rotation point is always where nums[i] > nums[i+1]
        - where i+1 is the correct min idx to start from
    
    once we have to located the fall-off point, we know the ending idx and we know the starting idx
    - now traverse from start -> end and make sure things are sorted!

    can we never find a fall-off point?
    - yes, possible because there are dupes 
    - [1,1,1] - no fall off point, we can return true in this case
    - [1,2,3,4,5] - has a fall off point; endIdx = n-1 and startIdx = 0
*/
func check(nums []int) bool {
    n := len(nums)
    end := -1   
    for i := 0; i < 2*n; i++ {
        if nums[i%n] > nums[(i+1)%n] {end = i%n; break}
    }
    if end == -1 {return true}
    // now loop from end+1 to end
    start := (end+1)%n
    for i := start; i%n != end; i++ {
        curr := nums[i%n]
        next := nums[(i+1)%n]
        if curr > next {return false}
    }
    return true
}