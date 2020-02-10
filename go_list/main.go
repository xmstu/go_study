package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	// 尾部添加元素
	l.PushBack("canon")

	// 头部添加
	l.PushFront(67)

	// 尾部添加后保存元素句柄
	element := l.PushBack("first")

	// 在first之后添加high
	l.InsertAfter("high", element)

	// 在first之前添加noon
	l.InsertBefore("noon", element)

	// 删除元素句柄
	l.Remove(element)

	// 遍历列表(双链表) 1.需要通过Front函数获取头元素;2.遍历时要判断当前节点是否为nil;3.不断调用Next函数获取下一个节点;
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
