// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

var PHONE_CHAR = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func TestLetterCombinations(t *testing.T) {
	n1 := "2"
	res1 := letterCombinations(n1)
	assert.Equal(t, []string{"a", "b", "c"}, res1)

	n2 := "23"
	res2 := letterCombinations(n2)
	println(res2)

	n3 := "234"
	res3 := letterCombinations(n3)
	println(res3)

}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	if len(digits) < 0 || len(digits) > 4 {
		fmt.Printf("digits nums must > 0 or < 4")
		return nil
	}
	bytes := []byte(digits)
	digit := []rune(string(bytes))
	nums := make([]int, len(digit))
	for i := range digit {
		num, err := strconv.Atoi(string(digit[i]))
		if err != nil {
			fmt.Printf("give not is num")
			return nil
		}
		if num > 9 || num < 2 {
			fmt.Printf("num > 10 or num < 2")
			return nil
		}
		nums[i] = num
	}
	return backTrack(nums, res, make([]string, 0), 0, len(nums))
}

// 回溯递归算法解决排列问题，回溯会把问题变成遍历树的枝干（边）
// num 当前层级输入的数值，如2、3、4
// res 结果
// track 记录的回溯排序的数值
// layer 回溯层级
// higher 总层高
func backTrack(nums []int, res []string, track []string, layer int, higher int) []string {
	// 递归终止条件
	if len(track) == higher {
		t := strings.Join(track, "")
		res = append(res, t)
		return res
	}
	// 如 2 -> abc
	letter := PHONE_CHAR[nums[layer]]
	bytes := []byte(letter)
	letterArray := []rune(string(bytes))
	for i := 0; i < len(letterArray); i++ {
		// 选择
		track := append(track, string(letterArray[i]))
		// 递归
		res = backTrack(nums, res, track, layer+1, higher)
		// 撤销选择
		track = track[0 : len(track)-1]
	}
	return res
}
