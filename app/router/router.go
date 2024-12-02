package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kanahiro/go-api/model"
)

type ModelPost struct {
	ID   string `json:"id" binding:"required"`
	Data string `json:"data" binding:"required"`
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		id := c.Query("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is required",
			})
			return
		}

		data, err := model.GetById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "get",
			"data":    data,
		})

	})

	r.POST("/", func(c *gin.Context) {
		var modelData ModelPost
		if err := c.ShouldBindJSON(&modelData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		newModel := &model.Model{
			ID:   modelData.ID,
			Data: modelData.Data,
		}

		newModel, err := model.PostNewData(newModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "post",
			"data":    newModel,
		})
	})
	return r
}
