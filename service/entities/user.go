package entities

type User struct {
	Key      int64  `xorm:"autoincr notnull pk 'key'"`
	UserName string `xorm:"username"`
	Password string `xorm:"password"`
	Email    string `xorm:"email"`
	Tel      string `xorm:"telephone"`
}

func NewUser(username string, password string) User {
	var user User
	user.UserName = username
	user.Password = password
	user.Email = ""
	user.Tel = ""
	return user
}
