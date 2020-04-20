package Git

// 父结构体，表示人
type Person struct {
	Name      string
	Email     string
}

// 子结构体，表示记录（人）
type Record struct {
	Person
	Time       uint64
}

func NewRecord(name string, email string, time uint64) *Record {
	return &Record{Person{name, email}, time}
}