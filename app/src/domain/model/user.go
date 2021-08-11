package model

import "time"

type UserModel struct {
	Id           int64     `json:"Id"`
	UserName     string    `json:"UserName"`
	UserEmail    string    `json:"UserEmail"`
	UserPassWord string    `json:"UserPassWord"`
	CreatedAt    time.Time `json:"Create_at"`
	UpdatedAt    time.Time `json:"Update_at"`
}
