/*
    To make a palindrome, we need 2 SAME EXACT CHARACTERS
    
    approach: freqMap
    - We can build out a freqMap { $char: count }
    - Then since we know we can need 2 count of same exact characters - this also means if the count is 4 , 6,8 ( any even num ) - we can take them all
    - So we can loop over the freqMap
    - and if the char count is even, take them all and add to our len var. delete this key from map
    - if the count is not even, than we can take max even from this count ( count-1 on a odd count will give us the highest even count we can get )
    - Then we will add the even count from odd count and update the count of this odd count char with whatever is left after taking even out. ( count-evenCountThatWeTook )
    - Finally once we have looped over the entire freqMap;
    - if our map still has any data ( i.e len(freqMap) != 0 )
    - We can atleast take 1 character and add to our len ( remember we are not forming the string but just the longest possible len of a palindrome string )
    - Why can we take 1 character and add that to our len count?
    - Because a string is still a palindrome if we can 1 character in the middle.
        - For example: input; abccccdd
        - dccccd - but we can add either a or b in the middle and this will still be a palindrome
        - dccaccd or dccbccd
    
    Potential follow up: print the string that's being formed 
    
*/

// func longestPalindrome(s string) int {
//     freqMap := map[byte]int{}
//     for i := 0; i < len(s); i++ {
//         freqMap[s[i]]++
//     }
    
//     maxLen := 0
//     for char, count := range freqMap {
//         if count % 2 == 0 {
//             maxLen+=count
//             delete(freqMap, char)
//         } else {
//             evenCount := count-1
//             maxLen += evenCount
//             freqMap[char] = count-evenCount
//         }
//     }
    
//     if len(freqMap) != 0 {maxLen++}
//     return maxLen
// }


func longestPalindrome(s string) int {
    seen := map[byte]bool{}
    maxLen := 0
    d := new(dq)
    for i := 0; i < len(s); i++ {
        _, ok := seen[s[i]]
        if ok {
            d.appendToEnd(string(s[i]))
            d.prepend(string(s[i]))
            maxLen += 2
            delete(seen, s[i])
        } else {
            seen[s[i]]=false
        }
    }
    if len(seen) != 0 {
        for k, _ := range seen{
            d.insertInMiddle(string(k))
            break
        }
        
        maxLen++
    }
    fmt.Println(d.valsToString())
    return maxLen
}


type listnode struct {
    val string
    next *listnode
}
type dq struct {
    head *listnode
    tail *listnode
    size int
}


func (d *dq) prepend(val string) {
    newNode := &listnode{val: val}
    if d.head == nil {
        d.head = newNode
        d.tail = newNode
        d.size = 1
        return
    }
    newNode.next = d.head
    d.head = newNode
    d.size++
}   

func (d *dq) appendToEnd(val string) {
    newNode := &listnode{val: val}
    if d.head == nil {
        d.head = newNode
        d.tail = newNode
        d.size = 1
        return
    }
    d.tail.next = newNode
    d.tail = newNode
    d.size++
}

func (d *dq) insertInMiddle(val string) {
    newNode := &listnode{val: val}
    if d.head == nil {
        d.appendToEnd(val)
        return
    }
    mid := d.size/2
    currPos := 0
    curr := d.head
    var prev *listnode
    for currPos != mid {
        prev = curr
        curr = curr.next
        currPos++
    }
    
    if prev != nil {
        prev.next = newNode
    }
    newNode.next = curr
}

func (d *dq) valsToString() string {
    str := new(strings.Builder)
    curr := d.head
    for curr != nil {
        str.WriteString(curr.val)
        curr = curr.next
    }
    return str.String()
}