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

type RegisterModel struct {
	UserName     string `json:"UserName"`
	UserEmail    string `json:"UserEmail" binding:"required"`
	UserPassWord string `json:"UserPassWord" binding:"required"`
}

type LoginModel struct {
	UserName     string `json:"UserName"`
	UserPassWord string `json:"UserPassWord" binding:"required"`
}

type UserBlog struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Title    string `json:"title"`
	Context  string `json:"context"`
}
