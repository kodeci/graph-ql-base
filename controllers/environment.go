package controllers

import (
	"ichabod/models"

	"github.com/gin-gonic/gin"
)

//EnvironmentController ...
type EnvironmentController struct{}

var environmentModel = new(models.EnvironmentModel)

// //Create ...
// func (ctrl EnvironmentModel) Create(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	var articleForm forms.ArticleForm

// 	if c.BindJSON(&articleForm) != nil {
// 		c.JSON(406, gin.H{"message": "Invalid form", "form": articleForm})
// 		c.Abort()
// 		return
// 	}

// 	articleID, err := environmentModel.Create(userID, articleForm)

// 	if articleID > 0 && err != nil {
// 		c.JSON(406, gin.H{"message": "Article could not be created", "error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Article created", "id": articleID})
// }

//All ...
func (ctrl EnvironmentController) All(c *gin.Context) {

	// data, err := environmentModel.All(userID)

	// if err != nil {
	// 	c.JSON(406, gin.H{"Message": "Could not get the articles", "error": err.Error()})
	// 	c.Abort()
	// 	return
	// }

	c.JSON(200, gin.H{"data": "get all results"})
}

// //One ...
// func (ctrl ArticleController) One(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	id := c.Param("id")

// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

// 		data, err := environmentModel.One(userID, id)
// 		if err != nil {
// 			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"data": data})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter"})
// 	}
// }

// //Update ...
// func (ctrl ArticleController) Update(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	id := c.Param("id")
// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

// 		var articleForm forms.ArticleForm

// 		if c.BindJSON(&articleForm) != nil {
// 			c.JSON(406, gin.H{"message": "Invalid parameters", "form": articleForm})
// 			c.Abort()
// 			return
// 		}

// 		err := environmentModel.Update(userID, id, articleForm)
// 		if err != nil {
// 			c.JSON(406, gin.H{"Message": "Article could not be updated", "error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"message": "Article updated"})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
// 	}
// }

// //Delete ...
// func (ctrl ArticleController) Delete(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	id := c.Param("id")
// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

// 		err := environmentModel.Delete(userID, id)
// 		if err != nil {
// 			c.JSON(406, gin.H{"Message": "Article could not be deleted", "error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"message": "Article deleted"})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter"})
// 	}
// }
