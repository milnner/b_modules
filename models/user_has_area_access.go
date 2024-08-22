package models

type UserHasAreaAccess struct {
	User User `json:"user"`
	Area Area `json:"area"`
}
