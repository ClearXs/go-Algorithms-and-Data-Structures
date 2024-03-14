package heap

// 堆用数组表示，但其可以用完全二叉树来表达
// 对于任意元素下标i，其子结点的元素是2i，2i-1，父结点的元素是[2/i]
// 对于数组下标为0的一定是其极值

type SortSymbolism = string

const (
	Asc  SortSymbolism = "asc"
	Desc SortSymbolism = "desc"
)

type Heap[T interface{}] struct {
	items      []*T
	sort       SortSymbolism
	comparator func(t1 *T, t2 *T) int
}

// NewHeap 创建完全二叉堆，值得注意的是该堆的的内置数组是从1开始的
func NewHeap[T interface{}](sort SortSymbolism, comparator func(t1 *T, t2 *T) int) *Heap[T] {
	return &Heap[T]{items: make([]*T, 1), sort: sort, comparator: comparator}
}

func NewMaximumHeap[T interface{}](comparator func(t1 *T, t2 *T) int) *Heap[T] {
	return NewHeap[T](Desc, comparator)
}

func NewMinimumHeap[T interface{}](comparator func(t1 *T, t2 *T) int) *Heap[T] {
	return NewHeap[T](Asc, comparator)
}

// InsertArray 插入数组
func (heap *Heap[T]) InsertArray(array []*T) {
	for _, ele := range array {
		heap.Insert(ele)
	}
}

// Insert 向堆中插入元素，需满足堆序性
func (heap *Heap[T]) Insert(t *T) {
	// 创建blank hole
	heap.items = append(heap.items, nil)
	holeIndex := heap.Size()
	// 初始化
	if holeIndex == 1 {
		heap.items[1] = t
		return
	}
	heap.siftUp(t, holeIndex)
}

// siftUp 堆的上滤，从堆尾部，向上寻找适合的位置进行插入
func (heap *Heap[T]) siftUp(temp *T, holeIndex int) {
	// 要么blank hole处于root，要么找到适合的root位置
	// percolate up
	for holeIndex > 1 {
		parentIndex := holeIndex / 2
		parent := heap.items[parentIndex]
		swap := heap.compareSwap(parent, temp)
		if swap {
			heap.items[holeIndex] = parent
			heap.items[parentIndex] = temp
			holeIndex = parentIndex
		} else {
			heap.items[holeIndex] = temp
			holeIndex = 0
		}
	}
}

// DeleteExtreme 删除极值
func (heap *Heap[T]) DeleteExtreme() *T {
	if heap.Size() == 0 {
		return nil
	}
	// 创建root blank hole
	extreme := heap.items[1]
	heap.items[1] = nil
	tail := heap.items[heap.Size()]
	heap.siftDown(tail, 1)
	// 去除末尾元素
	heap.items = heap.items[:heap.Size()]
	return extreme
}

// siftDown 堆的下滤，从上开始往下，给定元素，判断是否需要进行替换
func (heap *Heap[T]) siftDown(temp *T, holeIndex int) {
	// 基于完全二叉堆的特性，其left child index = 2 * i，right child index = 2 * i + 1

	// 1.比较两个child，判断哪一个child需要进行替换
	// 2.比较child与temp，如果child替换，则按照child index再次进行遍历
	// 3.循环结束的条件是该hole node没有child
	for holeIndex*2 < heap.Size() {
		childIndex := holeIndex * 2
		// 比较左右结点，寻找更适合替换hole的结点index
		rightSwap := heap.compareSwap(heap.items[childIndex], heap.items[childIndex+1])
		if rightSwap {
			childIndex += 1
		}
		// 比较与temp的结点
		tempSwap := heap.compareSwap(heap.items[childIndex], temp)
		if tempSwap {
			break
		} else {
			heap.items[holeIndex] = heap.items[childIndex]
			holeIndex = childIndex
		}
	}
	heap.items[holeIndex] = temp
}

func (heap *Heap[T]) Size() int {
	return len(heap.items) - 1
}

func (heap *Heap[T]) Peek(index int) *T {
	if index == 0 || index > len(heap.items)-1 {
		return nil
	}
	return heap.items[index]
}

// compare 基于给定的排序方式，可比较器进行比较，如果为true，说明需要交换位置，为false则不需要
func (heap *Heap[T]) compareSwap(t1 *T, t2 *T) bool {
	comparator := heap.comparator
	if comparator == nil {
		panic(comparator)
	}
	sort := heap.sort
	compareValue := comparator(t1, t2)
	if sort == Asc {
		if compareValue >= 0 {
			return true
		}
		return false
	} else if sort == Desc {
		if compareValue >= 0 {
			return false
		}
		return true
	}
	panic(comparator)
}
