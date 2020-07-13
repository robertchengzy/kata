GC 算法有四种:

引用计数（reference counting）
标记-清除（mark & sweep）
节点复制（Copying Garbage Collection）
分代收集（Generational Garbage Collection）。

go的对象大小定义:

大对象是大于32KB的.
小对象16KB到32KB的.
Tiny对象指大小在1Byte到16Byte之间并且不包含指针的对象.

三色标记的一个明显好处是能够让用户程序和 mark 并发的进行.