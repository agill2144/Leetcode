// func findAnagrams(s string, p string) []int {
//     if len(p) > len(s) {
// 		return nil
// 	}
// 	charFreq := map[string]int{}
// 	for _, ele := range p {
// 		charFreq[string(ele)]++
// 	}
// 	windowStart := 0
// 	totalMatches := 0
// 	out := []int{}
// 	for windowEnd := 0; windowEnd < len(s); windowEnd++{
// 		char := string(s[windowEnd])
// 		_, inPattern := charFreq[char]
// 		if inPattern {
// 			charFreq[char]--
// 			if val := charFreq[char]; val == 0 {
// 				totalMatches++
// 			}
// 		}
// 		if totalMatches == len(charFreq) {
// 			out = append(out, windowStart)
// 		}
// 		if windowEnd-windowStart+1 >= len(p) {
// 			leftChar := string(s[windowStart])
// 			windowStart++

// 			if val, inPatternMap := charFreq[leftChar]; inPatternMap {
// 				if val == 0 {
// 					totalMatches--
// 				}
// 				charFreq[leftChar]++
// 			}
// 		}
// 	}
// 	return out
// }

/*
Fixed window size - len(p) == windowSize
What defines a valid window == when size == windowSize
What defines whether we have found p in current window?
- we could keep track of all chars we have seen so far in this window
- and check with the characters in p string?
- how do we easily check whether current characters in current window are all in p ?
- keep current characters and their count in a windowMap and compare with characters in p map?
- if they all cross out then the windowStart is 1 answer
When does our windowSlide?
- When we do/not have all characters matching with pMap
What happens to our current window map when sliding the window
- decrement the freq count of the char leaving our window.

*/
// suboptimal brute forced 
// func findAnagrams(s, p string) []int {
// 	if len(p) > len(s) {
// 		return nil
// 	}
// 	pMap := map[string]int{}
// 	for _, ele := range p {
// 		pMap[string(ele)]++
// 	}
// 	left := 0
// 	windowMap := map[string]int{}
// 	out := []int{}
// 	for right := 0; right < len(s); right++ {
// 		char := string(s[right])
// 		windowMap[char]++
// 		windowSize := right-left+1
// 		if windowSize == len(p) {
//             // fmt.Println("windowMap: ", windowMap)
//             // fmt.Println("pMap: ", pMap)

//             countMatches := 0
// 			for k, v := range windowMap {
// 				if times, ok := pMap[k]; ok && times == v {
// 					countMatches++
// 				}
// 			}
//             // fmt.Println("CountMatches: ", countMatches)
// 			if countMatches == len(pMap) {
// 				out = append(out, left)
// 			}
// 			windowMap[string(s[left])]--
// 			if val, _ := windowMap[string(s[left])]; val == 0 {
// 				delete(windowMap, string(s[left]))
// 			}
// 			left++
// 		}
// 	}
// 	return out
// }

/*
    Basic Sliding window template
    
    left := 0
    while rightIdx < len(nums) {
        
        perform-calculations
        ( could be creating a window state in list, map)
        
        if windowSize == fixedSize {
            get answer from our calcuations first
            save answer whererev if we have one 
            
            undo the calculation for the item going out of the window
            ( remove,decrement,increment window state thats being maintained in a list/map)
            ( or something mathmatical - like undo the add - so minus, or undo the multiply - so divide etc)
            
            slide the window forward ( left++ )
        }
    
    }
    
    

*/
func findAnagrams(s, p string) []int {
    if len(p) > len(s) { return nil }
    pMap := map[string]int{}
    for _, ele := range p {
        pMap[string(ele)]++
    }
    
    left := 0
    out := []int{}
    count := len(pMap)
    for right := 0; right < len(s); right++ {
        // perform calc
        // for this we have to decrement the count of char in our pMap if found
        char := string(s[right])
        _, inPattern := pMap[char]
        if inPattern {
            pMap[char]--
            if val := pMap[char]; val == 0 {
                count--
            }
        }
        if right-left+1 == len(p) {
            // get the answer first
            if count == 0 {
                out = append(out, left)
            }
            
            // undo calculations
            leftChar := string(s[left])
            if charCount, ok := pMap[leftChar]; ok {
                if charCount == 0 {
                    count++
                }
                pMap[leftChar]++
            }
            left++
        }
    }
    return out
}
