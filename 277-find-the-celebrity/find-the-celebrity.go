/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        indegrees := make([]int, n)
        outdegrees := make([]int, n)
        for i := 0; i < n; i++ {
            // if a knows b; this is a->b ( a outdegrees++ ; b indegress++ )
            // its possible that a does not b; but b knows a or both is true
            // therefore we have to check in both direction
            // if b knows a; this is b->a (a indegrees++; b outdegrees++)
            for j := i; j < n; j++ {
                if i == j {continue}
                a := i; b := j
                if knows(a,b) {
                    indegrees[b]++
                    outdegrees[a]++
                }
                if knows(b, a) {
                    indegrees[a]++
                    outdegrees[b]++
                }
            }
        }
        // definition of a celebrity is that 
        // all the other n - 1 people know the celebrity
        // but the celebrity does not know any of them.
        // if indegrees[x] == n-1 && outdegrees[x] == 0 {this is celebrity}
        // fmt.Println(indegrees)
        // fmt.Println(outdegrees)
        for i := 0; i < n; i++ {
            if indegrees[i] == n-1 && outdegrees[i] == 0 {return i}
        }
        return -1
    }
}