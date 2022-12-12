package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(
		optimizeRoutes(
			[][]int{
				{1, 2000}, {2, 2000},
				{3, 3500}, {4, 4500}, {5, 6000}},
			[][]int{
				{1, 500}, {2, 1000},
				{3, 2500}, {4, 3000},
				{5, 3500}, {6, 5000}},
			7500))
	fmt.Println(optimizeRoutes([][]int{{1, 1000}, {2, 3000}, {3, 4000}}, [][]int{{1, 5000}, {2, 3000}}, 5000))

}

//	let m == smaller and n == bigger
// time: o(mlogn) -- if we do not count sorting foward and backward arrays
// space: o(1)
func optimizeRoutes(forward [][]int, backward [][]int, maxLimit int) []int {

	// time: o(flogf)
	sort.Slice(forward, func(i, j int) bool {
		return forward[i][1] <= forward[j][1]
	})

	// time: o(blogb)
	sort.Slice(backward, func(i, j int) bool {
		return backward[i][1] <= backward[j][1]
	})

	smaller := forward
	larger := backward
	if len(backward) < len(smaller) {
		smaller = backward
		larger = forward
	}
	closest := 0
	result := []int{}

	// time: o(min(len(forward),len(backward)) * log(max(len(forward),len(backward))
	for i := 0; i < len(smaller); i++ {
		rem := maxLimit - smaller[i][1]
		idx := binarySearchFloorOfTarget(larger, rem)
		if idx != -1 {
			dist := smaller[i][1] + larger[idx][1]
			if dist > closest {
				closest = dist
				result = []int{smaller[i][0], larger[idx][0]}
			}
			if closest == maxLimit {
				break // not going to find anything better than this, stop searching
			}
		}
	}
	return result
}

func binarySearchFloorOfTarget(nums [][]int, target int) int {
	left := 0
	right := len(nums) - 1
	idx := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid][1] <= target {
			if nums[mid][1] == target {
				return mid
			}
			idx = mid
			left = mid + 1
		} else if nums[mid][1] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}
