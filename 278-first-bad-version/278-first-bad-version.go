/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
/*
    approach : Binary search
    - if at mid, the version is good
        - this means the version before this is also good
        - which means no need to search left anymore
    - if at mid, the version is bad
        - this maybe a new bad version we discovered, so save its version
        - but this does not necessarily mean, its the FIRST bad version
        - however this does mean, any new version to the right (i.e ; next versions ) are also bad
        - so no need to search right
        - but continue searching left
    - until left and right ptrs have crossed each other
    
    time: o(logn)
    space: o(1)
*/
func firstBadVersion(n int) int {
    
    left := 1
    right := n
    first := -1
    for left <= right {
        mid := left + (right-left)/2
        if isBadVersion(mid) {
            first = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return first
}