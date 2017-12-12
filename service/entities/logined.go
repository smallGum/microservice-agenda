package entities

// LoginInfo record logined information
type LoginInfo struct {
	Key      int64  `xorm:"autoincr notnull pk 'key'"`
	UserName string `xorm:"username"`
}
