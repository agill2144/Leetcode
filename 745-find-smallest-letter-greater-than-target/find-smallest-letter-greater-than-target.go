func nextGreatestLetter(letters []byte, target byte) byte {
    // find the left most on right side of target ( but not equal to target )
    left := 0
    right := len(letters)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        midVal := int(letters[mid]-'a')
        targetVal := int(target-'a')
        if midVal > targetVal {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    // "If such a character does not exist, return the first character in letters"
    if ans == -1 {return letters[0]}
    return letters[ans]
}