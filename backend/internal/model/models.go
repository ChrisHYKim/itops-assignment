package model

import "time"

// User Info 구성
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// user Error 출력
type Issue struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"descripton"`
	Status      string    `json:"status"`
	User        *User     `json:"user,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Users assigiment 구성
var Users = []User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}

// 사용자 검토 후, 정보를 저장합니다.
func GetUserByID(id uint) *User {
	for _, user := range Users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

// Issue 상태 반환 string valid Function
func IsValidStatus(status string) bool {
	switch status {
	case "PENDING", "IN_PROGRESS", "COMPLETED", "CANCELLED":
		return true
	default:
		return false
	}
}
