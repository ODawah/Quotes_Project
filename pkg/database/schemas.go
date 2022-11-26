package database

type Author struct {
	Id int
	Name string
}

type Quote struct {
	Id int
	Text string
	AuthorName string
	AuthorId int
}

type AuthorQuotes struct {
	Auth Author
	Quotes []Quote
}
