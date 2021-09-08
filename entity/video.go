package entity

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Age int8 `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" binding:"required,email"`
}

type Video struct {
	Title string `json:"title" binding:"min=2,max=10" validate:"is-cool"` // min 2 char, max 10 char
	Description string `json:"description" binding:"max=20"`
	URL string `json:"url" binding:"required,url"`
	Author Person `json:"author" binding:"required"`
}
