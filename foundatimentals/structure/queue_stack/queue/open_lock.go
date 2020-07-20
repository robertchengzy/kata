package queue

/*
打开转盘锁
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
示例 1:
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:
输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:
输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:
输入: deadends = ["0000"], target = "8888"
输出：-1
提示：
死亡列表 deadends 的长度范围为 [1, 500]。
目标数字 target 不会在 deadends 之中。
每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
*/

/*我们可以将 0000 到 9999 这 10000 状态看成图上的 10000 个节点，两个节点之间存在一条边，当且仅当这两个节点对应的状态只有 1 位不同，
且不同的那位相差 1（包括 0 和 9 也相差 1 的情况），并且这两个节点均不在数组 deadends 中。那么最终的答案即为 0000 到 target 的最短路径。

我们用广度优先搜索来找到最短路径，从 0000 开始搜索。对于每一个状态，它可以扩展到最多 8 个状态，即将它的第 i = 0, 1, 2, 3 位增加 1 或减少 1，
将这些状态中没有搜索过并且不在 deadends 中的状态全部加入到队列中，并继续进行搜索。注意 0000 本身有可能也在 deadends 中。
*/

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, v := range deadends {
		dead[v] = true
	} // 填充dead set
	if dead["0000"] {
		return -1
	} // 直接死锁
	if target == "0000" {
		return 0
	} // 出发即是终点，特殊

	// BFS --------------------------------------------------------------
	queue := make([]string, 0)         // 构造处理字符串队列
	queue = append(queue, "0000")      // 起点
	visited := make(map[string]uint16) // 已访问过的集合。由于总共只有一万个状态点，所以步数不可能需要更多，所以uint16足以表示
	visited["0000"] = 0

	var cur string
	var curSlice []byte
	var nexts [8]string
	var origin byte
	for len(queue) != 0 {
		cur = queue[0]         // 取出当前待处理的锁状态（或者说无向图的节点）
		queue = queue[1:]      // 出队
		curSlice = []byte(cur) // 转为切片

		// 获取当前状态下一步的所有（8个）可能状态
		for i := 0; i < 4; i++ { // 对每一位进行变动。
			origin = curSlice[i] // 备份下原始的字符
			// 正向转动转盘
			curSlice[i] = (curSlice[i]-'0'+1)%10 + '0' // '0'~'9'的字符减去'0' 变为整型，来和1作加减，外边再 + '0'又转为字符
			nexts[2*i] = string(curSlice)
			curSlice[i] = origin // 恢复原始状态
			// 反向转动转盘
			curSlice[i] = (curSlice[i]-'0'+9)%10 + '0' // 如果是-1会出现负数情况，不好处理。循环左移的技巧
			nexts[2*i+1] = string(curSlice)
			curSlice[i] = origin
		}

		// 遍历下一步的所有可能状态
		for _, next := range nexts {
			if _, ok := visited[next]; !ok && !dead[next] { // 没有访问过，也不是dead
				queue = append(queue, next)      // 入队
				visited[next] = visited[cur] + 1 // 步数增加
				if next == target {
					return int(visited[next])
				} // 如果到达目标，就返回最少需要的步数
			}
		}

	}

	return -1
}
