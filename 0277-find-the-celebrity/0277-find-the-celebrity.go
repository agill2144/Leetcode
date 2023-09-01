/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        potential := 0
        for i := 0; i < n; i++ {
            if knows(potential, i) {
                // if $potential knows i, that means $potential->i, so now i is the new possibility
                // a celeberity must NOT know anyone, if $potential knows i, $potential is no longer the right ans
                potential = i
            }
        }
        
        // now check if everyone else knows $potential
        for i := 0; i < n; i++ {
            if potential == i {continue}
            /*
                $potential knows someone else or
                someone else does not know $potential,
                we dont have an ans
            */
            if knows(potential, i) || !knows(i, potential) {return -1}
        }
        
        return potential

    }
}