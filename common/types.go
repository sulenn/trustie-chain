package common

const (
	// 区块哈希值长度
	BlockHashLength = 32
	// Git 哈希值长度
	GitHashLength = 40
)

//32 字节 Keccak256 哈希值
type BlockHash [BlockHashLength]byte

//40 字节 SHA-1 哈希值
type GitHash [GitHashLength]byte
