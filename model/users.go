package model

//PageInfo model
type PageInfo struct {
	Page  int
	Limit int
}

//Users model
type Users struct {
	TotalCount int
	Items      []*User
	PageInfo   *PageInfo
}
