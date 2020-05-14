package models

type Email struct {
	Email string `form:"email" json:"email" binding:"required"`
}
