package test

import (
	"fmt"
	"testing"
	"unsafe"
)

type I interface {
	Say()
}

type P struct {
}

func (P) Say() {
	fmt.Println(1)
}

func Test1(t *testing.T) {

	var _ I = (*P)(nil)

	//
	//var a, b uint
	//a = 1
	//b = 2
	//fmt.Println(a - b)   // 答案，结果会溢出，如果是32位系统，结果是2^32-1，如果是64位系统，结果2^64-1.

	//
	//
	//runtime.GOMAXPROCS(1)
	//var wg sync.WaitGroup
	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go func(n int) {
	//		defer wg.Done()
	//		fmt.Println(n)
	//	}(i)
	//}
	//fmt.Println(runtime.NumGoroutine())
	//wg.Wait()

	// slice 是指针类型， append 会扩容， 证明扩容后使用的是新的缓存空间
	//s := make([]int, 1, 2)
	//b := s
	//fmt.Println(s, b)  // [0] [0]
	//
	//s = append(s, 1)
	//s[0] = 1
	//fmt.Println(s, b) // [1, 2] [1] , 但是如果修改 append 的数量， 则 b 会变为 0， 代表已经不是同一个 slice 了

	// 实现一个 interface ，
	// 值接受 那么 struct 本身和指针都实现了接口
	// 指针接受，那么 只有 struct 的指针实现了接口
	//var a I
	//a = P{}
	//a = &P{}

}

type person struct {
	name string
	age  int
}

/**

文章详细:  https://segmentfault.com/a/1190000040858899

1 任何类型的指针值都可以转换为 Pointer。
2 Pointer 可以转换为任何类型的指针值。

3 uintptr 可以转换为 Pointer
4 Pointer 可以转换为 uintptr
*/

func TestUnsafePoint(t *testing.T) {

	n := person{"su", 18}

	namePointer := (*string)(unsafe.Pointer(&n))
	*namePointer = "q"

	agePointer := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + unsafe.Offsetof(n.age)))
	*agePointer = 19

	fmt.Println(n)
}
