package database

import (
	"reflect"
)

type DatabaseConn struct {
	Class ConString

	User ConString

	Content ConString

	Area ConString

	ImageActivity ConString

	TextActivity ConString

	OneQuestionNAnswerActivity ConString

	AnswerNToOneActivity ConString
}

func NewDatabaseConn() *DatabaseConn { return &DatabaseConn{} }

type ConString struct {
	Insert string
	Select string
	Delete string
	Update string
}

func SetRoot(c *DatabaseConn, value string) {
	v := reflect.ValueOf(c).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Struct {
			fieldPtr := field.Addr().Interface()

			for j := 0; j < reflect.ValueOf(fieldPtr).Elem().NumField(); j++ {
				subField := reflect.ValueOf(fieldPtr).Elem().Field(j)
				if subField.Kind() == reflect.String {
					subField.SetString(value)
				}
			}
		}
	}
}

func (c *ConString) SetInsert(value string) {
	c.Insert = value
}

func (c *ConString) SetSelect(value string) {
	c.Select = value
}

func (c *ConString) SetDelete(value string) {
	c.Delete = value
}

func (c *ConString) SetUpdate(value string) {
	c.Update = value
}

func (c *ConString) GetInsert() string {
	return c.Insert
}

func (c *ConString) GetSelect() string {
	return c.Select
}

func (c *ConString) GetDelete() string {
	return c.Delete
}

func (c *ConString) GetUpdate() string {
	return c.Update
}
