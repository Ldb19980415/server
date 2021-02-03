package sunjie


type CreateWeight struct {
	CurrentWeight string	`form:"currentWeight" binding:"required" `
}
