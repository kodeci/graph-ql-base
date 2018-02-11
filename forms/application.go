package forms

//ApplicationCreateForm ...
type ApplicationCreateForm struct {
	Title string `form:"title" json:"title" binding:"required"`
}
