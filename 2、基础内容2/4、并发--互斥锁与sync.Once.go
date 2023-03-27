package main

import (
	"fmt"
	"image"
	"sync"
)

var (
	x int64
	//wg sync.WaitGroup // 等待组
	m             sync.Mutex // 互斥锁
	loadIconsOnce sync.Once
	icons         map[string]image.Image
)

func add() { // add 对全局变量x执行5000次加1操作
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func test() { // 1-未使用互斥锁，数据竞争的示例
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
func add1() { // add 对全局变量x执行5000次加1操作
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}
func main() { //2- 互斥锁，数据竞争的示例
	wg.Add(2)

	go add1()
	go add1()

	wg.Wait()
	fmt.Println(x)
}

//	func loadIcons() { //3-不安全的并发
//		icons = map[string]image.Image{
//			"left":  loadIcon("left.png"),
//			"up":    loadIcon("up.png"),
//			"right": loadIcon("right.png"),
//			"down":  loadIcon("down.png"),
//		}
//	}
// func Icon(name string) image.Image { // Icon 被多个goroutine调用时不是并发安全的
//
//		if icons == nil {
//			loadIcons()
//		}
//		return icons[name]
//	}

// func loadIcons() { //3-1安全的并发
//
//		icons = map[string]image.Image{
//			"left":  loadIcon("left.png"),
//			"up":    loadIcon("up.png"),
//			"right": loadIcon("right.png"),
//			"down":  loadIcon("down.png"),
//		}
//	}
//
// func Icon(name string) image.Image {// Icon 是并发安全的
//
//		loadIconsOnce.Do(loadIcons) //先执行检查
//		return icons[name]
//	}
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton { //sync.Once实现的并发安全的单例模式
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
