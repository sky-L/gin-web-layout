package test

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os/exec"
	"testing"
	"time"
	"unsafe"
)

type PI interface {
	ChangeName()
}

type P1 struct {
	Name string
}

func (p *P1) ChangeName() {
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

func changeM(m map[int]int) {
	m[1] = 2
}

func String2Byte(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func TestSomeCode(t *testing.T) {

	b, err := execSync("./", "ls", "-l")
	fmt.Println(err, string(b))

	time.Sleep(3 * time.Second)

	//e := errors.WithStack(errors.New("错"))
	//
	//fmt.Printf("%+v", e)

	// 生成随机字符串
	//buf := make([]byte, 10)
	//_, _ = io.ReadFull(rand.Reader, buf)
	//secret := hex.EncodeToString(buf)
	//fmt.Println(secret)

	// map 会变， 因为 map 底层初始化返回的是指针
	//m := map[int]int{1: 1}
	//fmt.Println(m)
	//changeM(m)
	//fmt.Println(m)

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

	// 值接受者方法，默认实现指针接受者方法   反之不行
	//var p PI =  P1{"L"}
	//var p PI =  &P1{"L"}
	//p.ChangeName()
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

func execSync(pwd string, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = pwd

	buf := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	go io.Copy(buf, stdout)
	go io.Copy(bufErr, stderr)
	if err := cmd.Run(); err != nil {
		e := bufErr.String()
		return nil, errors.New(e)
	}
	return buf.Bytes(), nil
}
