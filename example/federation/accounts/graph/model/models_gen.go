// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Entity struct {
	FindUserByID *User `json:"findUserByID"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (User) Is_Entity() {}
