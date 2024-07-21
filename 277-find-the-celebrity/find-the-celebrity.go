/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        celeb := 0
        for i := 0; i < n; i++ {
            if i == celeb {continue}
            if knows(celeb, i) {
                celeb = i
            }
        }
        for i := 0; i < n; i++ {
            if i == celeb {continue}
            if !knows(i, celeb) || knows(celeb, i) {return -1}
        }
        return celeb
    }
}