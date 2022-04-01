package test

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)
import "github.com/spaolacci/murmur3"


var once sync.Once

func TestM(t *testing.T) {

	ret := map[uint32]int{}
	for i := 1; i < 100; i++ {
		si := strconv.Itoa(i)
		m := murmur3.Sum32([]byte("176002426" + si))
		ret[m] = 1
	}
	fmt.Println(ret)

}

func TestSnowflake(t *testing.T) {

	s := []int{}

	s = append(s, 1, 2, 3)
	fmt.Println(len(s), cap(s))

	//arr := []int{1,2,3}
	//
	//arr1 := arr[:0]
	//fmt.Println(cap(arr1), len(arr1))
	//
	//fmt.Println(append(arr[0:1], arr[2:]...))

	//s := 'a'
	//fmt.Println(int(s))
	//
	//fmt.Println(string(rune(65)))

	//n, _ := snowflake.NewNode(1)
	//
	//id := n.Generate()
	//
	//
	//fmt.Println(id, id.Step())

}

func TestOnce(t *testing.T)  {

	config := atomic.Value{}
	config.Load()

	for i := 1; i< 10; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
}