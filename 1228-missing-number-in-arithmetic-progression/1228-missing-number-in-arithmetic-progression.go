/*
    approach: binary search
    - binary search before numbers are sorted
    - first we need to find the pattern ( common difference )
        - is it +1 sequence, +2, +3 or some +n ?
    - we know that the 0th idx and last idx numbers are not changed
    - and there is a missing number ( therefore array len is 1 element shorter )
    - To find the common differnce
        - example; [5,7,11,13]
        - how many hops do we have to take to reach 13 in above array ?
        - 4 hops ( len of array )
        - so we need find a $diff that when added to 5, 4 times ( because hops ) should equal 13
        - 13 = 5 + ($diff x 4)
        - (13-5)/4 = $diff
        - which will give us the common diff
        - this translates to (arr[n-1]-arr[0])/n
    - now we have a common diff; the arithemtic progession
    - what should be the value at idx 1, if diff was 2
        - [5,x,9,11,13]
        - add $diff to 5 N times ( where N is the distance between 5 and x )
        - this N also happens to be the idx
        - so 5 + ($diff * 1)
        - 5 + (2*1) = 7
    - what should be the value at idx 3m if diff was 2
        - [5,7,9,x,13]
        - add $diff to 5 N times ( where N is the distance between 5 and x ; 3-0 = 3 )
        - this N also happens to be the idx of x
        - so 5 + ($diff * 3)
        - 5 + (2*3) = 11
    - so we now how to get the common diff
    - once we also have a common difference, we can tell which value is supposed to exist at a specific idx
    - now using the above 2, run a binary search
        - is mid the missing number?
            - if the expected value at mid is not the same value at mid idx
            - then this may be a potential answer ( expected value may be the missing number )
            - but it could also be that an interruption has happened before reaching this idx.
            - so save this expected value as a potential answer and check left more ( right = mid-1 )
            - if the expected value at mid is the same value at mid idx;
            - it means an interruption in arithmetic progression has not happened yet;
            - therefore no need to search left, search right ( left = mid + 1 )
    - finally return the expected missing value that we had been saving all this while.
  
  time : o(logn)
  space: o(1)
*/
func missingNumber(arr []int) int {
    n := len(arr)
    // if all numbers are the same [0,0,0] or [1,1,1] ; there is no missing number
    if arr[n-1] == arr[0] {return arr[0]}
    diff := (arr[n-1]-arr[0])/n
    left := 0
    right := n-1
    missing := -1
    for left <= right {
        mid := left+(right-left)/2
        expectedMid := arr[0] + ( mid * diff )
        if arr[mid] != expectedMid {
            missing = expectedMid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return missing
}
