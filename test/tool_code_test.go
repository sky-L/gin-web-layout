package test

import (
	"fmt"
	"testing"
)

func TestDelSlice(t *testing.T) {

	i := 1
	a := []int{1,3}
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	fmt.Println(a)


	// map 中间插入元素
	//m := make([]int, 0, 0)
	//m = []int{1, 3, 4}
	//m = append(m[:1], append([]int{2}, m[1:]...)...)
	//fmt.Println(m)

	// rand.Int()

	// 切片的属性 len 和 cap 是值，非指针，所以 append 要返回新的 slice

	// 字符串底层也是字节数组
	// 字符串支持 slice 操作

	// 1
	//s1 := []int{1}
	//for i := 0; i < 5; i++ {
	//	// 当 append 的次数超过切片容量的时候， 每次都会内存重新分配
	//	s1 = append(s1, i)
	//}

	// 2

	// 内存模型

	//  1 顺序一致性内存模型

}

func Inc() (v int) {
	defer func() { v++ }()
	return 42
}
