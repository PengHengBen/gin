package models

import "time"

type User struct {
	Id             int
	Name           string
	Gender         string
	Age            byte
	SubmissionDate time.Time
}

// 表示配置库的表名称
func (user User) TableName() string {
	return "user"
}
