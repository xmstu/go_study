package main

import (
	"fmt"
	"sync"
)

func syncMap() {
	var scene sync.Map

	// 保存键值对
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 根据键读取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}

func main() {
	syncMap()

	m := make(map[int]int)

	go func() {
		// 不停地对map进行写入
		for {
			m[1] = 1
		}
	}()

	// 开启一段并发代码
	go func() {
		// 不停地对map进行读取
		for {
			_ = m[1]
		}
	}()

	// 主线程无限循环，让并发程序在后台执行
	// 并发读写map会提前发现，并终止程序执行
	for {

	}

}
