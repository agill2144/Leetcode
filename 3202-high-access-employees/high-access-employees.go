/*
    approach: sort first, then use sliding window
    tc = o(nlogn) + o(2n) = o(nlogn) + o(n)
    sc = o(1)
*/
func findHighAccessEmployees(access_times [][]string) []string {
    sort.Slice(access_times, func(i, j int)bool{
        if access_times[i][0] == access_times[j][0] {
            return access_times[i][1] < access_times[j][1]
        }
        return access_times[i][0] < access_times[j][0]
    })
    
    out := []string{}
    i := 0
    for i < len(access_times) {
        name := access_times[i][0]
        left := i
        right := i
        for right < len(access_times) && access_times[right][0] == name {
            if right < len(access_times) && right-left+1 < 3 {right++; continue}
            for left <= right && timeDiffInMins(left, right, access_times) >= 60 {
                left++
            }
            if right-left+1 >= 3 {
                out = append(out, name)
                break
            }
        }
        i = right
        for i < len(access_times) && access_times[i][0] == name {i++}
    }
    return out
}

func timeDiffInMins(left, right int, times [][]string) int {
    leftTime := times[left][1]
    leftHH,_ := strconv.Atoi(leftTime[:2])
    leftMM,_ := strconv.Atoi(leftTime[2:])
    leftMins := (leftHH*60) + leftMM

    rightTime := times[right][1]
    rightHH,_ := strconv.Atoi(rightTime[:2])
    rightMM,_ := strconv.Atoi(rightTime[2:])
    rightMins := (rightHH*60) + rightMM
    return rightMins-leftMins
}
/*
    approach: 2 pass + hashmap + sorting
    - parse times into minutes
    - group each persons times in a list of times per person
        - { $name : [time1,time2, etc]}
    - then loop over the map (name, times)
    - if we do not have atleast 3 times, we can skip this person
    - now we need a window of size atleast 3
    - which can confirm that the difference between each element is < 60 mins
    - we can sort the array, so we only need to confirm the difference between 2 extreme ends
        - 0th and last elemet in our 3 size subarray
    - we dont need to check adj times, 
    - because, the largest possible diff in a sorted array
    - will always be the lastElement - firstElement
    - if our extreme end is satisfied ( right Time - left time < 60 mins )
    - the entire subarray is satisfied
    - if we have a subarray of size 3, save this name

    tc = o(n) + o(nlogn * n)
    sc = o(n)
*/
// func findHighAccessEmployees(access_times [][]string) []string {
//     parsedTimings := map[string][]int{}
//     for i := 0; i < len(access_times); i++ {
//         name := access_times[i][0]
//         t := access_times[i][1]
//         hh, _ := strconv.Atoi(t[:2])
//         mm, _ := strconv.Atoi(t[2:])
//         mins := (hh*60) + mm
//         parsedTimings[name] = append(parsedTimings[name], mins)
//     }
//     out := []string{}
//     for name, times := range parsedTimings {
//         if len(times) < 3 {continue}
//         sort.Ints(times)
//         left := 0
//         for i := 0; i < len(times); i++ {
//             // get a window of size 3 first
//             if i-left+1 < 3 {continue}
//             // when we have a window of size 3
//             // then make sure its the 2 entreme ends
//             // are within the limit (highest time, lowest time)
//             // if they are not < 60 mins, shrink window from left
//             for left <= i && times[i] >= times[left]+60 {
//                 left++
//             }
//             // we dont need to check adj times, 
//             // because, the largest possible diff in a sorted array
//             // will always be the lastElement - firstElement
//             // if our extreme end is satisfied ( right Time - left time < 60 mins )
//             // the entire subarray is satisfied
//             // if we have a subarray of size 3, save this ans
//             if i-left+1 == 3 {out =append(out, name); break}
//         }
//     }
//     return out
// }