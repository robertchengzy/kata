package design

/*
设计哈希集合
不使用任何内建的哈希表库设计一个哈希集合
具体地说，你的设计应该包含以下的功能
	add(value)：向哈希集合中插入一个值。
	contains(value) ：返回哈希集合中是否存在这个值。
	remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
示例:
	MyHashSet hashSet = new MyHashSet();
	hashSet.add(1);
	hashSet.add(2);
	hashSet.contains(1);    // 返回 true
	hashSet.contains(3);    // 返回 false (未找到)
	hashSet.add(2);
	hashSet.contains(2);    // 返回 true
	hashSet.remove(2);
	hashSet.contains(2);    // 返回  false (已经被删除)
注意：
	所有的值都在 [0, 1000000]的范围内。
	操作的总数目在[1, 10000]范围内。
	不要使用内建的哈希集合库。
*/

type MyHashSet struct {
	set [][]int
}

const SET_MAX_LEN = 100000

func getIndex(key int) int {
	return key % SET_MAX_LEN
}

func (this *MyHashSet) getPos(key, index int) int {
	temp := this.set[index]
	if temp == nil {
		return -1
	}

	for i, data := range temp {
		if data == key {
			return i
		}
	}
	return -1
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{set: make([][]int, SET_MAX_LEN)}
}

func (this *MyHashSet) Add(key int) {
	index := getIndex(key)
	pos := this.getPos(key, index)
	if pos < 0 {
		this.set[index] = append(this.set[index], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	index := getIndex(key)
	pos := this.getPos(key, index)
	if pos >= 0 {
		this.set[index] = append(this.set[index][:pos], this.set[index][pos+1:]...)
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := getIndex(key)
	pos := this.getPos(key, index)
	return pos >= 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
