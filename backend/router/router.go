package router

import (
	v1 "github.com/coding-codes/api/v1"
	"github.com/coding-codes/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 如果请求的 method 不对，给出相应的回应
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
	})

	apiv1 := r.Group(utils.AppInfo.ApiBaseUrl)

	{
		apiv1.Static(utils.AppInfo.StaticBasePath, utils.AppInfo.UploadBasePath)
		apiv1.POST("/user/login", v1.Login)
		apiv1.POST("/user/logout", v1.Logout)
		apiv1.GET("/user/info", v1.GetUserInfo)
		apiv1.GET("/user/about", v1.GetUserAbout)
		apiv1.PATCH("/user/edit", v1.EditUserAbout)
		apiv1.POST("/upload", v1.UploadImageAvatar)

		// tags
		apiv1.GET("/tags", v1.GetAllTags)
		apiv1.POST("/tags", v1.CreateTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// category
		apiv1.GET("/categories", v1.GetAllCategory)
		apiv1.POST("/categories", v1.CreateCategory)
		apiv1.PUT("/categories/:id", v1.EditCategory)
		apiv1.DELETE("/categories/:id", v1.DeleteCategory)

		// soups
		apiv1.POST("/soups", v1.CreateSoup)
		apiv1.DELETE("/soups/:id", v1.DeleteSoup)
		apiv1.PUT("/soups/:id", v1.EditSoup)
		apiv1.GET("/soups", v1.GetAllSoups)
		apiv1.GET("/soup/random", v1.GetRandSoup)
	}
	return r
}
