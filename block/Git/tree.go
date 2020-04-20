package Git

import (
	"github.com/trustie-chain/common"
)

type Tree struct {
	//CurrentHash    common.GitHash
	Size           int
	Items          []*Item
}


type Item struct {
	Type           string  // 只有blob 和 tree 两种类型
	Hash           common.GitHash
	Name           string  // 文件或目录名
}

func NewTree(size int, items []*Item) *Tree {
	return &Tree{size,items}
}
