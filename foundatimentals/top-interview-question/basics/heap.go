package basics

import (
	"fmt"
)

// a := []int{2,7,26,25,19,17,1,90,3,36}
// [90 36 17 25 26 7 1 2 3 19]

// 堆
// O(NlogN)
func InitHeap(a []int) {
	length := len(a)
	heap := make([]int, length)
	for i := 0; i < length; i++ {
		// 插入后和父节点比较如果大于父节点则交换
		heap[i] = a[i]
		leaf := i + 1
		parent := leaf / 2
		for parent > 0 {
			if heap[leaf-1] > heap[parent-1] {
				heap[leaf-1], heap[parent-1] = heap[parent-1], heap[leaf-1]
			}
			leaf = parent
			parent = parent / 2
		}
	}

	fmt.Println(heap)
}

// O(N)
func InitHeapSimple(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		down(a, i, n)
	}
}

func down(a []int, i0, n int) bool {
	i := i0 // 父级
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // 左子树
		if j2 := j1 + 1; j2 < n && a[j2] > a[j1] {
			j = j2 // 右子树
		}
		if a[i] >= a[j] {
			break
		}

		a[i], a[j] = a[j], a[i]
		i = j
	}

	return i > i0
}
