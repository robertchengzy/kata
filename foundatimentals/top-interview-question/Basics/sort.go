package basics

import "fmt"

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
		for j := i + 1; j < len(s); j++ { //从后面挑选最小的
			if s[i] > s[j] { //拿定点，跟后面的比较，最小的就交换位置
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// 插入排序：从后面乱序中，依次取出一个然后插入到前面排序的位置，
func sortByInsert(s []int64) {
	for i := 1; i < len(s); i++ { //i控制后面乱序和前面顺序的分界点，i控制分界点的移动
		if s[i-1] > s[i] { //判断相邻的位置是否大小顺序正确，否则就要找到正确的位置
			for j := i - 1; j >= 0 && s[j] > s[j+1]; j-- { //j控制位置的移动
				s[j+1], s[j] = s[j], s[j+1] //以交换的方式做了简单的数据移动
			}
		}
	}
	fmt.Println(s)
}

// 快速排序
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

func quickSort(s []int64, low, high int64) { //start起始点，end终止点
	if low < high {
		flag := partition(s, low, high) //把切片一分为二，分别对两部分进行递归排序
		quickSort(s, low, flag-1)       //低的部分进行排序
		quickSort(s, flag+1, high)      //高的部分进行排序
	}

}

func partition(s []int64, low, high int64) int64 {
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
		fmt.Println(s)
	}

	return low //返回本次排序的能够确定最终位置的元素位置
}
