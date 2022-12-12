/*
    approach: try to build out all sequences with numbers that we can start a sequence from
    - First we need to make our elements in nums array searchable fast, therefore we will put them in a hashmap
    - Then while looping over the nums array, we will;
        - Check if the current number is a start of a sequence
        - current number would be start of a sequence IF there is not a number that comes before this number
            - check if our map contains currentNumber-1 
            - if it does, this means that this currentNumber is not the start of a sequence..
            - REMEMBER FOR EACH ELEMENT WE ARE CHECKING IF THIS NUMBER COULD BE THE START OF A SEQUENCE
            - IF IT IS, WE KEEP CHECKING ITS NEXT NUMBER AND SAVING LENGTH EACH TIME
            - if there is no prev number in our nums array of our currentNumber
            - this means this currentNumber is definetely start of a sequence
            - Now, from this number maintain a count (starting at 1 to include the start sequence number as part of count)
            - Keep adding 1 to this currentNumber and as long as our nums map has the next number , keep incrementing the count
            - Once we no longer have a next immediate number, our loop will break out and we will need to save our count ( max = math.max(max, count))
    
    
    - OR work backwards -- look for a number that will end a sequence and count back as long as you can.
    
    time: o(n)
    space: o(1)
*/
// approach: looking for a start of sequence and counting as far as we can go
func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    
    max := 0
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        // is this number start of a sequence?
        _, exists := set[num-1];
        // if there is no prev number than this number, than this number is the start of sequence
        // try to build the sequence with this number as long as possible by adding +1 to num and checking in set
        if !exists {
            count := 1
            next := num+1
            _, setContainsNext := set[next]
            for setContainsNext {
                count++
                next++
                _, setContainsNext = set[next]
            }
            if count > max {
                max = count
            }
        }
    }
    return max
}

// approach: looking for an end of sequence and counting back as far as we can go
// func longestConsecutive(nums []int) int {
//     set := map[int]struct{}{}
//     for i := 0; i < len(nums); i++ {
//         set[nums[i]] = struct{}{}
//     }
    
//     max := 0
//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         // is this number end of a sequence?
//         _, exists := set[num+1];
//         // if there is no next number than this number, than this number is the end of sequence
//         // try to build the sequence with this number as long as possible by going back 1 from current num and checking in set
//         if !exists {
//             count := 1
//             prev := num-1
//             _, setContainsPrev := set[prev]
//             for setContainsPrev {
//                 count++
//                 prev--
//                 _, setContainsPrev = set[prev]
//             }
//             if count > max {
//                 max = count
//             }
//         }
//     }
//     return max
// }