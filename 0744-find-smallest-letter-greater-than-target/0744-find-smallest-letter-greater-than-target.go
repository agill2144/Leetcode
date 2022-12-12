func nextGreatestLetter(letters []byte, target byte) byte {
    if target >= letters[len(letters)-1] {return letters[0]}
    left := 0
    right := len(letters)-1
    var nextCeil byte
    
    for left <= right {
        mid := left+(right-left)/2
        if letters[mid] > target {
            nextCeil = letters[mid]
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nextCeil
}