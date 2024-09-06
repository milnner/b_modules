package models

type UserHasClassAccess struct {
	User  User  `json:"user"`
	Class Class `json:"class"`
}
