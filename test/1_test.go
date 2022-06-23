package test

import (
	"testing"
)

//type PI interface {
//	ChangeName()
//}

type P1 struct {
	Name string
}

func (p P1) ChangeName() {
	p.Name = "k"
}

type iface struct {
	itab, data uintptr
}

type MyWrite struct {
}

func (m *MyWrite) Write(p []byte) (n int, err error) {
	return
}

func TestSomeCode(t *testing.T) {



	// 检查类型，是否实现了接口
	// var _ io.Writer = (*MyWrite)(nil)

	// interface 动态类型、动态值
	//var a interface{} = nil
	//
	//var b interface{} = (*int)(nil)
	//
	//x := 5
	//var c interface{} = &x
	//
	//ia := *(*iface)(unsafe.Pointer(&a))
	//
	//ib := *(*iface)(unsafe.Pointer(&b))
	//
	//ic := *(*iface)(unsafe.Pointer(&c))
	//
	//fmt.Println(ia, ib, ic)
	//
	//fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
	// interface 动态类型、动态值 end

	//
	////var p PI
	//p := &P1{"L"}
	//
	//p.ChangeName()
	//
	//fmt.Println(p)

	//s := make([]int, 0, 5)
	//s = append(s, 1)
	//
	//mys := myAppend(s)
	//
	//fmt.Println(cap(s), s, mys)

	//changeS(s)
	//fmt.Println(s)

}

func myAppend(s []int) []int {
	s[0] = 2
	b := append(s, 100)
	return b
}

func changeS(s []int) {
	for k := range s {
		s[k] = 2
	}
}
