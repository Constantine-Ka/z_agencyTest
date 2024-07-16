package db

//go:generate reform

// reform:News
type News struct {
	ID      int64  `reform:"Id,pk"`
	Title   string `reform:"Title"`
	Content string `reform:"Content"`
}

type NewsCategories struct {
	NewsId     int64 `reform:"NewsId"`
	CategoryId int64 `reform:"CategoryId"`
}
