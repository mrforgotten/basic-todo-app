package dto

type TodoCreate struct {
	Title string `binding:"required,min=3,max=100"`
}

type TodoUpdate struct {
	Title string `binding:"required,min=3,max=100"`
}
