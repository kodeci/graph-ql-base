package forms

//CreateForm ...
type CreateForm struct {
	Title string `form:"title" json:"title" binding:"required"`
}
