package Git

import "github.com/trustie-chain/common"

type Commit struct {
	Author         *Record
	Committer      *Record
	Tree           common.GitHash
	//CurrentHash    common.GitHash
	ParentHash     common.GitHash
	Size           int
	Message        string
}

func NewCommit(author *Record, committer *Record, tree common.GitHash, parentHash common.GitHash, size int, message string) *Commit {
	return &Commit{author,committer,tree,parentHash, size,message}
}

//func NewCommit(author *Record, committer *Record, tree common.GitHash, currentHash common.GitHash, parentHash common.GitHash, size int, message string) *Commit {
//	return &Commit{author,committer,tree,currentHash,parentHash, size,message}
//}