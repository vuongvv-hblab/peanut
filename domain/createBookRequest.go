package domain

type CreateBookReq struct {
	Name string `json:"name" binding:"required"`
}
