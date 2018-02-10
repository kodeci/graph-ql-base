package controllers

import (
	"ichabod/forms"
	"ichabod/models"

	"github.com/gin-gonic/gin"
)

//ApplicationController ...
type ApplicationController struct{}

var applicationModel = new(models.ApplicationModel)

//All ...
func (ctrl ApplicationController) All(c *gin.Context) {

	c.JSON(200, gin.H{"data": "get all applications results"})
}

//Create ...
func (ctrl ApplicationController) Create(c *gin.Context) {
	var createForm forms.CreateForm

	if c.BindJSON(&createForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": createForm})
		c.Abort()
		return
	}

	application, err := applicationModel.Create(createForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if application.ID > 0 {
		// session := sessions.Default(c)
		// session.Set("application_id", application.ID)
		// session.Set("application_title", application.Title)
		// session.Save()
		c.JSON(200, gin.H{"message": "Success create", "application": application})
	} else {
		c.JSON(406, gin.H{"message": "Could not create applicaiton", "error": err.Error()})
	}

}
