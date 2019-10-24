package resolver

import (
	"gql-ashadi/model"
)

//PageInfoResolver page info resolver
type PageInfoResolver struct {
	page *model.PageInfo
}

//Page get page
func (r PageInfoResolver) Page() int32 {
	return int32(r.page.Page)
}

//Limit get limit
func (r PageInfoResolver) Limit() int32 {
	return int32(r.page.Limit)
}

//NewPageInfoResolver create new pageinfo resolver
func NewPageInfoResolver(p *model.PageInfo) *PageInfoResolver {
	return &PageInfoResolver{page: p}
}

//UsersResolver users resolver
type UsersResolver struct {
	users *model.Users
}

//TotalCount get total count of record
func (r UsersResolver) TotalCount() int32 {
	return int32(r.users.TotalCount)
}

//PageInfo get page info
func (r UsersResolver) PageInfo() *PageInfoResolver {
	return NewPageInfoResolver(r.users.PageInfo)
}

//Users get users record
func (r UsersResolver) Items() *[]*UserResolver {
	resolvers := make([]*UserResolver, 0)
	for _, user := range r.users.Items {
		resolvers = append(resolvers, NewUserResolver(user))
	}
	return &resolvers
}

//NewUsersResolver create new instance of users resolver
func NewUsersResolver(u *model.Users) *UsersResolver {
	return &UsersResolver{users: u}
}
