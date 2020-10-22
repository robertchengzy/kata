package design

/*
设计哈希映射
不使用任何内建的哈希表库设计一个哈希映射
具体地说，你的设计应该包含以下的功能
	put(key, value)：向哈希映射中插入(键,值)的数值对。如果键对应的值已经存在，更新这个值。
	get(key)：返回给定的键所对应的值，如果映射中不包含这个键，返回-1。
	remove(key)：如果映射中存在这个键，删除这个数值对。
示例：
	MyHashMap hashMap = new MyHashMap();
	hashMap.put(1, 1);
	hashMap.put(2, 2);
	hashMap.get(1);            // 返回 1
	hashMap.get(3);            // 返回 -1 (未找到)
	hashMap.put(2, 1);         // 更新已有的值
	hashMap.get(2);            // 返回 1
	hashMap.remove(2);         // 删除键为2的数据
	hashMap.get(2);            // 返回 -1 (未找到)
注意：
	所有的值都在 [0, 1000000]的范围内。
	操作的总数目在[1, 10000]范围内。
	不要使用内建的哈希库。
*/

type MyHashMap struct {
	hMap [][]pair
}

type pair struct {
	key, value int
}

const MAP_MAX_LEN = 100000

func getMapIndex(key int) int {
	return key % MAP_MAX_LEN
}

func (this *MyHashMap) getPos(key, index int) int {
	temp := this.hMap[index]
	if temp == nil {
		return -1
	}

	for i, data := range temp {
		if data.key == key {
			return i
		}
	}
	return -1
}

/** Initialize your data structure here. */
func ConstructorX() MyHashMap {
	return MyHashMap{
		hMap: make([][]pair, MAP_MAX_LEN),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := getMapIndex(key)
	pos := this.getPos(key, index)
	if pos < 0 {
		data := pair{
			key:   key,
			value: value,
		}
		this.hMap[index] = append(this.hMap[index], data)
	} else {
		this.hMap[index][pos].value = value
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := getMapIndex(key)
	pos := this.getPos(key, index)
	if pos < 0 {
		return -1
	}
	return this.hMap[index][pos].value
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := getMapIndex(key)
	pos := this.getPos(key, index)
	if pos >= 0 {
		this.hMap[index] = append(this.hMap[index][:pos], this.hMap[index][pos+1:]...)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
