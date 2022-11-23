package model

type DemoCommentDTO struct {
	Id     int    `json:"id"`
	PostId int    `json:"post_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
