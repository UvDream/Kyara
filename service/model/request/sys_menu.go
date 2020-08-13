package request

import "blog-api/model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string
}

// Get role by id structure
type AuthorityIdInfo struct {
	AuthorityId string
}
