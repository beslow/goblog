package models

type Category struct {
	ID   int
	Name string
}

type CategoryWithNum struct {
	Category
	Num int64
}
