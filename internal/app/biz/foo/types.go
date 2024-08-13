package foo

type Foo struct {
	ID   int64  `json:"id" structs:"id"`
	Name string `json:"name" structs:"name"`
}
