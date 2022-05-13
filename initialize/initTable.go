package initialize

import "gin/models"

func InitTable() {
	user := models.User{}
	b := DB.Migrator().HasTable("user")
	if !b {
		DB.Migrator().CreateTable(&user)
	}
	user1 := models.User1{}
	bo := DB.Migrator().HasTable("user1")
	if !bo {
		DB.Migrator().CreateTable(&user1)
	}
}
