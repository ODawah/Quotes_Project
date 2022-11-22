package database

type Author struct {
	Id int
	Name string
}

type Quote struct {
	Id int
	Text string
	AuthorId string
}