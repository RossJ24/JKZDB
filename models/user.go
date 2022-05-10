package models

import "fmt"

type User struct {
	Email           string `json:"email"`
	Balance         int64  `json:"balance"`
	Age             int32  `json:"age"`
	AccountOpenedAt int64  `json:"account_opened_at"`
	LastUsed        int64  `json:"last_used"`
}

func (user *User) ToMap() map[string]string {
	rep := make(map[string]string)
	rep["email"] = user.Email
	rep["balance"] = fmt.Sprint(user.Balance)
	rep["age"] = fmt.Sprint(user.Age)
	rep["account_opened_at"] = fmt.Sprint(user.AccountOpenedAt)
	rep["last_used"] = fmt.Sprint(user.LastUsed)
	return rep
}
