package sort3

import "math/rand"

var N1 = []int{3, 5, 4, 1, 2, 6}
var N2 = []int{6, 5, 4, 3, 2, 1}
var N3 []int
var N4 = []int{6, 3, 2, 6, 7}

var EXPECT_N1 = []int{1, 2, 3, 4, 5, 6}
var EXPECT_N2 = []int{1, 2, 3, 4, 5, 6}
var EXPECT_N3 []int
var EXPECT_N4 = []int{2, 3, 6, 6, 7}

func Swap(nums []int, p int, q int) {
	swap := nums[q]
	nums[q] = nums[p]
	nums[p] = swap
}

func RandMoreNums(length int) []int {
	nums := make([]int, length)

	for j := 0; j < length; j++ {
		nums[j] = rand.Intn(length)
	}
	return nums

}
