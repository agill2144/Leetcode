// cancel out the valid pairs from input
// we only care about invalid ones
// specifically, we only care about openings that never closed
// if opening params count is x 
// the to close them we need x closing parans, yea?
// ]]]][[[[
// open = 4
// so we need 4 closing pairs to make the above input valid
// swap 0 and last idx; []]][[[]
// the above fixed 2 pairs!
// swap 1 and 3rd last idx; [] [] [] []
// again fixes 2 pairs
// therefore each swap fixes 2 pairs
// and if there are 4 pairs to fix, and each swap fixes 2
// then first swap fixes 2 out of 4; left with 2
// then second swap fixes another 2 out of 2; left with 0
// therefore openParansCount/2
// NOW THE EDGE CASE WITH ABOVE IS that even count vs odd count
// if we had even count of unbalanced parans, then 2 fixes per swap works!
// i.e openCount/2 gives us min number of swaps
// if the number of unbalanced open count was odd?
// ] ]][[ [
// open count = 3
// res = 3/2 = 1 -- this is saying we need 1 swap to fix
// lets swap last and 0th idx; [] ][ []
// fixes 2 pairs like we have seen
// but we need another swap!
// therefore if count is odd then we need +1 more swap
// open count = 4
// and we know each swap fixes 2 pairs ( drops open count by 2)
// open count / 2 = 2 swaps
// open count = 5
// and we know each swap fixes 2 pairs ( drops open count by 2)
// 1st swap, open count = 3 ( 5-2 )
// 2nd swap, open count = 1 ( 3-2 )
// see how we still have 1 open paran!
// thats why we need +1 more swap! when open count is odd
func minSwaps(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            count++
        } else {
            count--
            if count < 0 {count=0}
        }
    }
    
    if count % 2 == 0 {return count/2}
    return (count+1)/2
}