package models

import "time"

type User1 struct {
	Id             int
	Name           string
	Gender         string
	Age            byte
	SubmissionDate time.Time
}

// 表示配置库的表名称
func (user1 User1) TableName() string {
	return "user1"
}
