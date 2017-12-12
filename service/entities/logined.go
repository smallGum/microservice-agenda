package entities

import (
	"errors"
	"strconv"
)

// LoginInfo record logined information
type LoginInfo struct {
	Key      uint64 `xorm:"autoincr notnull pk 'key'"`
	UserName string `xorm:"username"`
}

// CheckKey check if the key of user is valid
func CheckKey(key string) (string, error) {
	var login LoginInfo
	k, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return "", err
	}
	got, err := agendaDB.Id(k).Get(&login)

	if !got {
		return "", errors.New("invalid key")
	}
	if err != nil {
		return "", err
	}
	return login.UserName, nil
}
