package forms

//EnvironmentCreateForm ...
type EnvironmentCreateForm struct {
	Title string `form:"title" json:"title" binding:"required"`
	Slug  string `form:"slug" json:"slug" binding:"max=100"`
}

//EnvironmentUpdateForm ...
type EnvironmentUpdateForm struct {
	Title string `form:"title" json:"title"`
	Key   string `form:"key" json:"key" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
}
