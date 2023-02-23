/*
    approach: linear scan
    - hIdx for an ith position is basically how many papers have >= citations[i]
    - example
        - [0,1,3,5,6]
        - i = 0; how many papers have >= 0 citations ; 5
        - i = 1; how many papers have >= 1 citations ; 4
        - i = 2; how many papers have >= 3 citations ; 3
        - i = 3; how many papers have >= 5 citations ; 2
        - i = 4; how many papers have >= 6 citations ; 1
        - for an ith position how many papers are to the right side of it ( including itself )
        - that is; len(citations)-i
    - Now we can get hIdx for each ith position we need to find where h >= h and n-h <= h is met
    - linear scan from left to right while calculating hIdx for each ith position
    - as soon as our condition is met, we can break and return the hIdx for that ith position
    
*/
func hIndex(citations []int) int {
    n := len(citations)
    for i := 0; i < n; i++ {
        h := n-i
        if h <= citations[i] {return h} 
    }
    return 0
}