func findJudge(n int, trust [][]int) int {
    outdegrees := make([]int, n+1) // counts number of votes given to others ( i.e num of outgoing edges )
    indegrees := make([]int, n+1) // counts number of votes recieved ( i.e num of incoming edges )
    for i := 0; i < len(trust); i++ {
        // a trusts b
        // a -> b
        a,b := trust[i][0], trust[i][1]
        outdegrees[a]++
        indegrees[b]++
    }
    for i := 1; i < len(indegrees); i++ {
        votesRec := indegrees[i]
        votesGiven := outdegrees[i]
        if votesGiven == 0 && votesRec == n-1 {return i}
    }
    return -1
}