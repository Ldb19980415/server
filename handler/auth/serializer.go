package auth


type loginSerializer struct {
	UserName string	`json:"username" binding:"required" `
	PassWord string	`json:"password" binding:"required" `
}


type Person struct {
	Name string `json:"name" binding:"required"` // json格式从name取值，并且该值为必须的
	Age  int    `json:"age" binding:"required,gt=20"` // json格式从age取值，并且该值为必须的，且必须大于20
}