package dao

import (
	"fmt"
)

// User represents a user in Firebase
type User struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// UserStringFormat is the format to describe a user
const UserStringFormat = "{id:%v name:%v email:%v}"

// String describe the object
func (user *User) String() string {
	return fmt.Sprintf(UserStringFormat,
		user.ID,
		user.Name,
		user.Email)
}
