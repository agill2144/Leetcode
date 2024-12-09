/*
    How to pick a random number betwee 0 and n ?
    - rand.Intn(n) ; this excludes n, if n needs to be included: rand.Intn(n+1)
    - The probability of each number from 0 to n is 1/n
        - meaning each val has a probability of 1 out of n times

    - we have to pick a random idx
    - but each idx has a different probability
    - this is defined by their weights

    approach: brute force
    - expanded each idx into w[i] times
    - so that their repetition represents their probability
       0 1 2      0 1 2 3 4 5
    - [2,1,3] -> [0,0,1,2,2,2]
    - then we can use rand.Intn(len(expandedIdxs))
    - each idx now repeats its weight times, therefore their probabilt has also increased
    
    tc:
    - constructor: o(n * sumOfAllWeights)
    - PickIndex: o(1)
    sc:
    - o(sumOfAllWeights) for expandedIdxs array

    approach: binary search
    - instead of allocating such a large expandedIdxs array
    - we can create a compressed version of the same size as weights array
    - each val in this compressed array represents a range of idxs that are possible for $idx
       0 1 2      0 1 2 3 4 5
    - [2,1,3] -> [0,0,1,2,2,2]
    - idx 0 has 2 elements; it covers idxs 0,1
    - idx 1 has 1 element; it covers idxs prevEndRange+1 to prevEndRange+1 = [2,2] 
    - idx 2 has 3 elements; it covers idxs prevEndRange+1 to prevEndRange+3 = [3,5]
    - this can be represented using cumulative sum
       0 1 2      0 1 2
    - [2,1,3] -> [2,3,6]
    - idx 0 val 2 means, range ends at idx 2(excluding) from expanded array
    - idx 1 val 3 means, range ends at idx 3(excluding) from expanded array
    - idx 2 val 6 means, range ends at idx 6(excluding) from expanded array
    - now we will pick a random integer value using n = last element of cumulative sum
    - rIdx(randomIdx) = 5 ( from 0 to 6, excluding 6 )
    - now find where does idx 5 shows up in compressedRanges
    - idx 0 ends at idx 2) and 5 > 2
    - idx 1 ends at idx 3) and 5 > 3
    - idx 3 ends at idx 6) and 5 < 6, so potential answer = 3
    - we have no more idxs left.
    - cumulative sum is a sorted array
    - and we essentially need to look for the left most on right side of rIdx
    - the left-most of right side of 5 is 6 and its at idx 2, therefore 2 is our answer

    tc:
    - constructor: o(n)
    - N = total sum
    - pickIdx(): o(logN)
    sc:
    - o(n) in constructor
*/

type Solution struct {
    compressedRanges []int    
}


func Constructor(w []int) Solution {
    compressedRanges := make([]int, len(w))
    for i := 0; i < len(w); i++ {
        compressedRanges[i] = w[i]
        if i-1 >= 0 {compressedRanges[i] += compressedRanges[i-1]}
    }
    return Solution{compressedRanges}
}


func (this *Solution) PickIndex() int {
    n := len(this.compressedRanges)
    total := this.compressedRanges[n-1]
    rIdx := rand.Intn(total)
    left := 0
    right := n-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if rIdx >= this.compressedRanges[mid] {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */