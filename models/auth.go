package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitAuth() {
	db.AutoMigrate(&Auth{})
	db.Create(&Auth{
		Username: "admin",
		Password: "123456",
	})
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
