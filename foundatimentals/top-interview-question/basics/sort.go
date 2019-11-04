package basics

/*
	1.基于比较的排序算法:
		BUB - 冒泡排序,
		SEL - 选择排序,
		INS - 插入排序, (O(N^2) 基于比较的排序算法)
		MER - 归并排序 (递归实现),
		QUI - 快速排序 (O（N log N）递归实现),
		R-Q - 随机快速排序 (递归实现). 分而治之
	2.不基于比较的排序算法:
		COU - 计数排序,
		RAD - 基数排序.(O(N)

	计数排序
	假设：如果要排序的项目是小范围的整数，我们可以计算每个整数（在这个小范围内）的出现频率，然后通过循环该小范围来按排序顺序输出项目。
	在上面的示例数组中，对所有整数都在[1..9]内的示例数组尝试Counting Sort，因此我们只需要计算整数1出现的次数，出现整数2，...，出现整数9，然后遍历 如果频率 [y] = x，则 1至9 打印出整数 y 的 x 个副本。
	时间复杂度为O（N）来计算频率，O（N + k）以排序顺序输出结果，其中 k 是输入的整数范围，在本例中为9-1 + 1 = 9。 计数排序（Counting Sort）的时间复杂度为O（N + k），如果 k 很小，那么它就是O（N）。
	由于内存限制，当 k 相对较大时，我们将无法执行计数排序（Counting Sort）的计数部分，因为我们需要存储那些 k 个整数出现的次数。

	基数排序
	假设：如果要排序的项目是大范围但小数位的整数，我们可以将计数排序（Counting Sort）思想与基数排序（Radix Sort）结合起来，以实现线性时间复杂度。
	在基数排序中，我们将每个项目排序为一个 w 数字串（如果需要，我们填充小于w数字的前几个零的整数）。
	对于最低有效位（最右边）到最高有效位（最左边），我们通过 N 个项目并将它们按照活动数字放到10个队列中（每个数字[0..9]一个），就好像 一个修改的计数排序，因为这保留了稳定性。 然后我们再次重新连接组，以便进行后续迭代。
	请注意，我们只执行O（w×（N + k））次迭代。
*/

// 冒泡排序思想：从前到后，依次两两比较，两层循环，一层控制比较趟数，一层控制前后两两比较
func sortByBubble(s []int64) {
	for i := 0; i < len(s); i++ { //i控制比较趟数
		for j := 0; j < len(s)-i-1; j++ { //j控制从前到后，两两比较
			if s[j] > s[j+1] { //两两比较
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// 选择排序思想：假定位前面是即将要存有序的位置，后面是无序的位置。从没有排序的序列中，选择一个最小的，依次插入到前面的有序位置的后边。 ` 从乱序中找到目标 `
func sortBySelect(s []int64) {
	for i := 0; i < len(s); i++ { //控制比较的趟数,守住不动的定点
		min := i
		for j := i + 1; j < len(s); j++ { //从后面挑选最小的
			if s[min] > s[j] { //拿定点，跟后面的比较，最小的就交换位置
				min = j
			}
			s[i], s[min] = s[min], s[i]
		}
	}
}

// 插入排序：从后面乱序中，依次取出一个然后插入到前面排序的位置，
func sortByInsert(s []int64) {
	for i := 1; i < len(s); i++ { //i控制后面乱序和前面顺序的分界点，i控制分界点的移动
		if s[i-1] > s[i] { //判断相邻的位置是否大小顺序正确，否则就要找到正确的位置
			for j := i - 1; j >= 0; j-- { //j控制位置的移动
				if s[j] > s[j+1] {
					s[j+1], s[j] = s[j], s[j+1] //以交换的方式做了简单的数据移动
				} else {
					break
				}
			}
		}
	}
}

// 快速排序
/*
	划分步骤：
	选择一个项目 p（称为枢轴点）然后将 a[i..j] 的项目分为三部分：a [i..m-1]，a [m] 和 a[m + 1..j]。
	a [i..m-1]（可能为空）
	包含小于 p 的项目。 a [m] 是枢轴点 p，例如：指数 m 是已排序数组 a 的排序顺序中 p 的正确位置。
	a [m + 1..j]（可能为空）包含大于或等于 p 的项目。
	然后，递归地对这两部分进行排序
*/
func quickSort(a []int64, low, high int64) { // start起始点，end终止点
	if low < high {
		flag := partition(a, low, high) //把切片一分为二，分别对两部分进行递归排序 O(N)
		// a[low..high] ~> a[low..m–1], pivot, a[m+1..high]
		quickSort(a, low, flag-1) // 递归地将左子阵列排序
		// a[m] = pivot 在分区后就被排序好了
		quickSort(a, flag+1, high) // 然后将右子阵列排序
	}
}

func partition(a []int64, low, high int64) int64 {
	p := a[low] // p 是枢纽
	m := low    // S1 和 S2 一开始是空的
	for k := low + 1; k <= high; k++ {
		if a[k] < p {
			m++
			a[k], a[m] = a[m], a[k]
		} // 注意：在情况1的时候我们什么不做: a[k] >= p
	}
	a[low], a[m] = a[m], a[low]
	return m
}

/*
	排序算法主要思路：
	1.假定起始位置位哨兵
	2.写一个方法，作用用来确定中介的位置，并且进行同时的排序功能，返回中介的位置
	3.递归方式分别调用比中介点小的部分，和比中介点大的部分，一分为二的思想
	4.在求中介值时，始终明确low<high，使用一个temp变量存储哨兵的值
	5.从后往前，逐个遍历比较，遇到小于哨兵的就暂停，且进行交换
	6.从前往后，逐个遍历比较，遇到大于哨兵的就暂停，且进行交换
	7.直到low和high相遇，停止循环，返回low的位置，即下次的一分为二时使用的分界点
*/
func partition2(s []int64, low, high int64) int64 {
	//分别控制两个点，一个从前往后遍历，一个从后往前遍历
	//假设我们每次将序列中的第一个元素作为定位排序的目标
	tempValue := s[low] //哨兵
	for low < high {    //当两边相遇时，结束本趟比较，直到low和high相遇时本趟排序结束
		for s[high] > tempValue && low < high { //从后往前遍历，找比哨兵小的数
			high--
		}
		tempValue, s[high] = s[high], tempValue //遇到比哨兵小的数则暂停，进行交换

		for s[low] < tempValue && low < high { //然后，从前往后遍历，找比哨兵大的数
			low++
		}
		s[low], tempValue = tempValue, s[low] //遇到比哨兵大的，就暂停，进行交换
	}

	return low //返回本次排序的能够确定最终位置的元素位置
}

// 自顶向下的归并排序
// 归并排序非常适合分类非常大量的输入
func merge(a []int, low, mid, high int) {
	// subarray1 = a[low..mid], subarray2 = a[mid+1..high], both sorted
	length := high - low + 1
	tmp := make([]int, length)
	left, right, tmpIdx := low, mid+1, 0
	// 归并
	for left <= mid && right <= high {
		if a[left] <= a[right] {
			tmp[tmpIdx] = a[left]
			left++
		} else {
			tmp[tmpIdx] = a[right]
			right++
		}
		tmpIdx++
	}
	// leftover, if any
	for left <= mid {
		tmp[tmpIdx] = a[left]
		tmpIdx++
		left++
	}
	// leftover, if any
	for right <= high {
		tmp[tmpIdx] = a[right]
		tmpIdx++
		right++
	}
	// copy back
	for k := 0; k < length; k++ {
		a[low+k] = tmp[k]
	}
}

func mergeSort(a []int, low, high int) {
	// 要排序的数组是 a[low..high]
	if low < high { // base case: low >= high (0 or 1 item)
		mid := (low + high) / 2
		mergeSort(a, low, mid)    // 分成一半
		mergeSort(a, mid+1, high) // 递归地将它们排序
		merge(a, low, mid, high)  // 解决: 归并子程序
	}
}

// 自底向上归并排序
func mergeSortDownToUp(s []int) {
	n := len(s)
	for sz := 1; sz < n; sz *= 2 {
		for i := 0; i < n-sz; i += 2 * sz {
			merge(s, i, i+sz-1, min(i+2*sz-1, n-1))
		}
	}
}

func min(i, j int) int {
	if j < i {
		return j
	}
	return i
}

// 堆排序 O(N*logN)
func HeapSort(a []int) {
	// 建堆
	InitHeapSimple(a)
	length := len(a)
	heapSize := length - 1
	for i := length - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapSize = heapSize - 1
		MaxHeapify(a, heapSize, 0)
	}
}

// 计数排序 O(N)
func CountingSort(a []int, k int) {
	length := len(a)
	b := make([]int, length)
	c := make([]int, k)
	for i := 0; i < length; i++ {
		c[a[i]] = c[a[i]] + 1
	}
	for j := 0; j < length; j++ {
		c[j] = c[j] + c[j-1]
	}
	for m := length - 1; m >= 0; m++ {
		b[c[a[m]]] = a[m]
		c[a[m]] = c[a[m]] - 1
	}
}

// 基数排序 O(N)
func RadixSort(a []int) {

}

// 桶排序 O(N)
func BucketSort(a []int) {

}
