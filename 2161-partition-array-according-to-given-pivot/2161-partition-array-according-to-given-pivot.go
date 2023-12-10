func pivotArray(nums []int, pivot int) []int {
    out := []int{}
    equal := 0
    greater := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] < pivot {
            out = append(out, nums[i])
        } else if nums[i] == pivot {
            equal++
        } else if nums[i] > pivot {
            greater = append(greater, nums[i])
        }
    }
    for equal > 0 {
        out = append(out, pivot)
        equal--
    }
    if len(greater) > 0 {out = append(out, greater...)}
    return out
}
// func pivotArray(nums []int, pivot int) []int {
//     m := map[string][]int{"less": []int{}, "equal": []int{}, "greater": []int{}}
//     for i := 0; i < len(nums); i++ {
//         if nums[i] < pivot {
//             m["less"] = append(m["less"], nums[i])
//         } else if nums[i] == pivot {
//             m["equal"] = append(m["equal"], nums[i])
//         } else if nums[i] > pivot {
//             m["greater"] = append(m["greater"], nums[i])
//         }
//     }
//     out := []int{}
//     if len(m["less"]) > 0 {out = append(out, m["less"]...)}
//     if len(m["equal"]) > 0 {out = append(out, m["equal"]...)}
//     if len(m["greater"]) > 0 {out = append(out, m["greater"]...)}
//     return out
// }

// func pivotArray(nums []int, pivot int) []int {
//     out := []int{}
//     countEq := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] < pivot {
//             out = append(out, nums[i])
//         } else if nums[i] == pivot {
//             countEq++
//         }
//     }
//     for countEq > 0 {
//         out = append(out, pivot)
//         countEq--
//     }
    
//     for i := 0; i < len(nums); i++ {
//         if nums[i] > pivot {
//             out = append(out, nums[i])
//         }
//     }
//     return out
// }