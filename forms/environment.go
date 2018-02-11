package forms

//EnvironmentCreateForm ...
type EnvironmentCreateForm struct {
	Title string `form:"title" json:"title" binding:"required"`
	Slug  string `form:"slug" json:"slug" binding:"max=100"`
}
