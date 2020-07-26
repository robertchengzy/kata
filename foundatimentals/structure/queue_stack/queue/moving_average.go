package queue

/*
数据流中的移动平均值
题目描述:
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算其所有整数的移动平均值。
示例:
MovingAverage m = new MovingAverage(3);
m.next(1) = 1
m.next(10) = (1 + 10) / 2
m.next(3) = (1 + 10 + 3) / 3
m.next(5) = (10 + 3 + 5) / 3
*/

type MovingAverage struct {
	currentWindow []int
	previousSum   int
	maxSize       int
}

func (this *MovingAverage) Constructor() *MovingAverage {
	return &MovingAverage{
		currentWindow: nil,
		previousSum:   0,
		maxSize:       3,
	}
}

func (this *MovingAverage) Next(num int) int {
	if len(this.currentWindow) == this.maxSize {
		this.previousSum -= this.currentWindow[0]
		this.currentWindow = this.currentWindow[1:]
	}
	this.currentWindow = append(this.currentWindow, num)
	return this.previousSum / len(this.currentWindow)
}
