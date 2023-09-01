/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        a := 0
        for i := 0; i < n; i++ {
            if knows(a, i) {
                // if a knows b, that means a->b, so now b is the new possibility
                // a celeberity must NOT know anyone, if a know b, a is no longer the right ans
                a = i
            }
        }
        
        // now check if everyone else know a
        for i := 0; i < n; i++ {
            if i == a {continue}
            if knows(a, i) || !knows(i, a) {return -1}
        }
        
        return a

    }
}