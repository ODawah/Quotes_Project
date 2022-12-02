package database

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Quote struct {
	Id         int    `json:"id"`
	Text       string `json:"text"`
	AuthorName string `json:"author_name"`
	AuthorId   int    `json:"author_id"`
}

type AuthorQuotes struct {
	Auth   Author  `json:"auth"`
	Quotes []Quote `json:"quotes"`
}

type SearchQuote struct {
	Text string `json:"text"`
}
