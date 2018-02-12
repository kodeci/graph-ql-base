package controllers

import (
	"ichabod/forms"
	"ichabod/models"
	"strconv"

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
	var createForm forms.ApplicationCreateForm

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
		c.JSON(200, gin.H{"message": "Success create", "application": application})
	} else {
		c.JSON(406, gin.H{"message": "Could not create applicaiton", "error": err.Error()})
	}

	return
}

//One ...
func (ctrl ApplicationController) One(c *gin.Context) {
	applicationID := c.Param("appId")

	if id, err := strconv.ParseInt(applicationID, 10, 64); err == nil {

		data, err := applicationModel.One(id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Application not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}

	return
}
