package models

import "time"

type User struct {
	Id             int
	Name           string
	Gender         string
	Age            byte
	SubmissionDate time.Time
}

func (user User) TableName() string {
	return "user"
}
