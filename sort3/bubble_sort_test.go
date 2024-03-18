package sort3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, BubbleSort(N1))
	assert.Equal(t, EXPECT_N2, BubbleSort(N2))
	assert.Equal(t, EXPECT_N3, BubbleSort(N3))
	assert.Equal(t, EXPECT_N4, BubbleSort(N4))
}
