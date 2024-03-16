package leetcode

import (
	"data-structure/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 127. 单词接龙
// 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
//
// 每一对相邻的单词只差一个字母。
// 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
// sk == endWord
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordLength := len(wordList)
	if wordLength == 0 {
		return 0
	}
	if endWord != wordList[wordLength-1] {
		return 0
	}
	wordList = append([]string{beginWord}, wordList...)
	// breadth first search
	// initialization
	q := queue.NewQueue[int]()
	start := 0
	q.Push(&start)
	path := 0
	marked := make([]bool, len(wordList))

	for !q.Empty() {
		wordIndex := q.Pop()
		if marked[*wordIndex] {
			continue
		}
		marked[*wordIndex] = true
		word := wordList[*wordIndex]
		if word == endWord {
			return path + 1
		}
		// digraph adjoining table
		adj := relevant(*wordIndex, word, wordList)
		if len(adj) > 0 {
			path += 1
		}
		for _, wordIndex := range adj {
			q.Push(&wordIndex)
		}
	}
	return path
}

func relevant(start int, origin string, wordList []string) []int {
	adj := make([]int, 0)
	for i := start + 1; i < len(wordList); i++ {
		target := wordList[i]
		relevance := completeWithRelevance(origin, target)
		if relevance {
			adj = append(adj, i)
		}
	}
	return adj
}

func completeWithRelevance(origin string, target string) bool {
	if len(origin) != len(target) {
		return false
	}
	diffCount := 0
	for i := 0; i < len(origin); i++ {
		if origin[i] != target[i] {
			diffCount += 1
		}
		if diffCount > 1 {
			return false
		}
	}
	return true
}

func TestLadderLength(t *testing.T) {
	w1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	l1 := ladderLength("hit", "cog", w1)
	assert.Equal(t, 5, l1)

	w2 := []string{"hot", "dot", "dog", "lot", "log"}
	l2 := ladderLength("hit", "cog", w2)
	assert.Equal(t, 0, l2)
}
