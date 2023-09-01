/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        degrees := make([]int, n)
        for i := 0; i < n; i++ {
            a := i
            for j := 0; j < n; j++ {
                if i == j {continue}
                b := j
                if knows(a, b) {
                    degrees[a]--
                    degrees[b]++
                }
            }
        }
        for i := 0; i < len(degrees); i++ {
            if degrees[i] == n-1 {return i} 
        }
        return -1
    }
}