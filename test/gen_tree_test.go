package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Menu 菜单
type Menu struct {
	ID       int
	ParentID int
	Name     string
	Children []Menu
}

// TreeList 菜单
type TreeList struct {
	ID       int
	ParentID int
	Name     string
	Children []TreeList
}

// FormMenu 格式化菜单
func FormMenu(list []Menu, pid int) (formMenu []Menu) {
	for _, val := range list {
		if val.ParentID == pid {
			if pid == 0 {
				// 顶层
				formMenu = append(formMenu, val)
			} else {
				var children []Menu
				child := val
				children = append(children, child)
			}
		}
	}
	return
}

// GetMenu 获取菜单
func GetMenu(menuList []Menu, pid int) []TreeList {
	treeList := []TreeList{}
	for _, v := range menuList {
		if v.ParentID == pid {
			child := GetMenu(menuList, v.ID)
			node := TreeList{
				ID:       v.ID,
				Name:     v.Name,
				ParentID: v.ParentID,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}

func Test2(t *testing.T) {
	myMenu := []Menu{
		{ID: 1, ParentID: 0, Name: "用户管理"},
		{ID: 2, ParentID: 1, Name: "会员管理"},
		{ID: 3, ParentID: 2, Name: "权限管理"},
		{ID: 4, ParentID: 3, Name: "会员管理"},
		{ID: 5, ParentID: 4, Name: "黑名单"},
		{ID: 6, ParentID: 5, Name: "会员列表"},
	}

	list := GetMenu(myMenu, 0)
	b, _ := json.Marshal(list)
	fmt.Printf(string(b))
}
