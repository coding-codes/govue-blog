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
		apiv1.POST("/user/login", v1.Login)
		apiv1.POST("/user/logout", v1.Logout)
		apiv1.GET("/user/info", v1.GetUserInfo)
	}
	return r
}
