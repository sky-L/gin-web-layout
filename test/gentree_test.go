package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Node struct {
	Id       int
	Pid      int
	Children []*Node
}

func TestGentree(t *testing.T) {
	nodeList := []Node{
		{
			Id:       1,
			Pid:      0,
			Children: nil,
		},
		{
			Id:       2,
			Pid:      1,
			Children: nil,
		},
		{
			Id:       3,
			Pid:      2,
			Children: nil,
		},
	}

	items := make([]*Node, 0, len(nodeList))

	for _, v := range nodeList {
		items[v.Id] = &v
	}

	for _, item := range items {
		//  $items[$item['pid']]['children'][] = &$items[$item['id']];
		if v := items[item.Pid]; v != nil {
			items[item.Pid].Children = append(items[item.Pid].Children, items[item.Id])
		}
	}

	b, _ := json.Marshal(items)
	fmt.Println(string(b))
}

/**

$items = array_column($productCategoryList, null, 'id');


$a = [
	[
		'id' => 1,

	]
];
$tree = [];
foreach ($items as $item) {
    if (isset($items[$item['pid']])) {
        $items[$item['pid']]['children'][] = &$items[$item['id']];
    } else {
        $tree[] = &$items[$item['id']];
    }
}

*/

func TestGentree2(t *testing.T) {
	nodeList := []Node{
		{
			Id:       1,
			Pid:      0,
			Children: nil,
		},
		{
			Id:       2,
			Pid:      1,
			Children: nil,
		},
		{
			Id:       3,
			Pid:      2,
			Children: nil,
		},
	}

	items := map[int]Node{}

	for _, v := range nodeList {
		items[v.Id] = v
	}


	for _, item := range items {
		if _, ok := items[item.Pid]; ok {
			// $items[$item['pid']]['children'][] = &$items[$item['id']];

			tmp := items[item.Id]

			c := items[item.Pid]
			c.Children = append(c.Children, &tmp)

			items[item.Pid] = c
		}
	}

	//a, _ := json.Marshal(items)
	//
	//fmt.Println(string(a))
}

