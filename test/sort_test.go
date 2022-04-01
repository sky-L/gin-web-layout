package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"testing"
)

func TestMapSort(t *testing.T) {
	// Map ksort 的思路， 先去除 key， 然后 sort.Strings(), 再遍历取值
	l := map[string]string{
		"name": "xx",
		"age":  "xx",
	}
	kList := make([]string, 0, len(l))
	for k := range l {
		kList = append(kList, k)
	}
	sort.Strings(kList)
	signStr := ""
	for _, v := range kList {
		signStr += v + "=" + l[v] + "&"
	}
	signStr = signStr[0 : len(signStr)-1]

	md5Map := md5.New()
	md5Map.Write([]byte(signStr))
	s := hex.EncodeToString(md5Map.Sum(nil))

	s2 := fmt.Sprintf("%X", md5.Sum([]byte(signStr)))
	s3 := fmt.Sprintf("%d", md5.Sum([]byte(signStr)))

	fmt.Println(s, s2, s3)
}

func TestPersonNameSort(t *testing.T) {
	p := []Person{{Name: "ba", Age: 1}, {Name: "ab", Age: 2}}

	//	sort.Sort(ByeStructKey{p})
	// doc https://learnku.com/articles/38269
	//sort.SliceStable(p, func(i, j int) bool {
	//	return p[i].Name < p[j].Name
	//})

	sort.Slice(p, func(i, j int) bool {
		return p[i].Name < p[j].Name
	})

	// fmt.Scanln()
	fmt.Println(p)
}

type Person struct {
	Name string
	Age  int
}

type ByeStructKey struct {
	Person []Person
}

func (s ByeStructKey) Len() int {
	return len(s.Person)
}
func (s ByeStructKey) Swap(i, j int) {
	s.Person[i], s.Person[j] = s.Person[j], s.Person[i]
}
func (s ByeStructKey) Less(i, j int) bool {
	return s.Person[i].Name < s.Person[j].Name
}
