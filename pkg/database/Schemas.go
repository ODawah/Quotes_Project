package database

type author struct {
	id int
	name string
}

type quote struct {
	id int
	text string
	authorId string
}