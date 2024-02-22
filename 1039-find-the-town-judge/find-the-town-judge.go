func findJudge(n int, trust [][]int) int {
    outdegrees := make([]int, n+1) // counts giving trusts/vote to others
    indegrees := make([]int, n+1) // counts number of votes recieved
    for i := 0; i < len(trust); i++ {
        // a trusts b
        a,b := trust[i][0], trust[i][1]
        indegrees[b]++
        outdegrees[a]++
    }
    for i := 1; i < len(indegrees); i++ {
        votesRec := indegrees[i]
        votesGiven := outdegrees[i]
        if votesGiven == 0 && votesRec == n-1 {return i}
    }
    return -1
}