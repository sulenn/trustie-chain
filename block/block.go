package block

import (
	"github.com/trustie-chain/common"
	"github.com/trustie-chain/block/Git"
)

type Header struct {
	CurrentHash common.BlockHash
	ParentHash  common.BlockHash
	CommitHash  common.BlockHash
	Time        uint64
}

type Body struct {
	Commits []*Git.Commit
}