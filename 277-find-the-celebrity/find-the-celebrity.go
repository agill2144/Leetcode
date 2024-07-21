// assume a celebrity at first; $x
// now check if knows(x, y)
// if x knows y ; x -> y
// this means x is not a good candidate for sure,
// because for x to be celebrity it must have 0 outgoing edges
// in other words; x should not have known y if x was a true celeb
// but y can be, therefore now update our potential canditate to y
// and continue our evaluation
// then take another pass to confirm that our $potentialCandidate is truly a celeb
// meaning $potentialCandidate should not know ith person and this ith person must know $potentialCandidate
// if the above is not true, return -1
// if we did not return -1, we have confirmed that $potentialCandidate is our celebrity
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        // 1. eliminate people that cannot be celeb

        // assume starting celeb is person 0
        celeb := 0
        for i := 0; i < n; i++ {
            if i == celeb {continue}
            // if celeb knows someone else
            // celeb is not a valid celeb
            // but since this invalid celeb knows i
            // i could be a potential celeb
            // therefore we eliminate celeb and update it with someone else instead
            // celeb -> i is true
            // celeb is not a true celebrity; but i might be
            if knows(celeb, i) {celeb = i}
        }

        // 2. confirm our candidate
        // everyone should know celeb
        // and celeb should not know anyone
        for i := 0; i < n; i++ {
            if i == celeb {continue}
            if knows(celeb, i) || !knows(i, celeb) {return -1}
        }
        return celeb
    }
}
/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
// func solution(knows func(a int, b int) bool) func(n int) int {
//     return func(n int) int {
//         indegrees := make([]int, n)
//         outdegrees := make([]int, n)
//         for i := 0; i < n; i++ {
//             // if a knows b; this is a->b ( a outdegrees++ ; b indegress++ )
//             // its possible that a does not b; but b knows a or both is true
//             // therefore we have to check in both direction
//             // if b knows a; this is b->a (a indegrees++; b outdegrees++)
//             for j := i; j < n; j++ {
//                 if i == j {continue}
//                 a := i; b := j
//                 if knows(a, b) {
//                     indegrees[b]++
//                     outdegrees[a]++
//                 }
//                 if knows(b, a) {
//                     indegrees[a]++
//                     outdegrees[b]++
//                 }
//             }
//         }
//         // definition of a celebrity is that 
//         // all the other n - 1 people know the celebrity
//         // but the celebrity does not know any of them.
//         // if indegrees[x] == n-1 && outdegrees[x] == 0 {this is celebrity}
//         // fmt.Println(indegrees)
//         // fmt.Println(outdegrees)
//         for i := 0; i < n; i++ {
//             if indegrees[i] == n-1 && outdegrees[i] == 0 {return i}
//         }
//         return -1
//     }
// }