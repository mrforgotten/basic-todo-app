package dto

type AuthorCreateInput struct {
	Name string `binding:"required,min=3,max=200"`
}

type AuthorUpdateInput struct {
	Name string `binding:"required,min=3,max=200"`
}
