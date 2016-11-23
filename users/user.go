package users

import (
	"time"
)

// 構造体のフィールドは仮です。
type User struct {
	Id         int64     `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Status     string    `db:"status"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

func (u *User) isActivate() bool {
	return u.Status == "activete"
}
