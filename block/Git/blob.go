package Git

type Blob struct {
	//CurrentHash    common.GitHash
	Size           int
	Content        string
}

func NewBlob(size int, content string) *Blob {
	return &Blob{size, content}
}
