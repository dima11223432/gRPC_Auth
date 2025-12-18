package models

type User struct {
	ID       int
	email    string
	PassHash []byte
}
