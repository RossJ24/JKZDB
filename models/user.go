package models

type User struct {
	Email           string `json:"email"`
	Balance         int64  `json:"balance"`
	Age             int32  `json:"age"`
	AccountOpenedAt int64  `json:"account_opened_At"`
	LastUsed        int64  `json:"last_used"`
}

func (user *User) ToMap() map[string]string {
	rep := make(map[string]string)
	rep["email"] = user.Email
	rep["balance"] = string(user.Balance)
	rep["age"] = string(user.Age)
	rep["account_opened_at"] = string(user.AccountOpenedAt)
	rep["last_used"] = string(user.LastUsed)
}
