package forms

//EnvironmentCreateForm ...
type EnvironmentCreateForm struct {
	Title string `form:"title" json:"title" binding:"required"`
}
