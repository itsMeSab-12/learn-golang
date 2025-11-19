/*
Given an integer array nums, return true
if any value appears more than once in the array,
otherwise return false.
*/

package arrays

func ContainsDuplicate(nums []int) bool {
	hash := make(map[int]bool)
	for _, v := range nums {
		if exists := hash[v]; exists {
			return true
		}
		hash[v] = true
	}
	return false
}
