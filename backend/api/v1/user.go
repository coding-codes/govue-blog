package v1

import (
	"github.com/coding-codes/service"
	"github.com/coding-codes/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	data := make(map[string]interface{})
	user := service.User{}

	// 把 form 表单数据和后台对应的字段绑定
	if e := c.ShouldBindJSON(&user); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}

	// 验证登录信息
	isExist := user.CheckAuth()

	if isExist {
		// 登录信息验证成功后，生成 token 并返回
		token, e := user.GenToken()
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40004, nil, e))
			return
		}
		data["token"] = token
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	// 登录信息验证失败返回
	c.JSON(http.StatusUnauthorized, utils.GenResponse(40001, nil, nil))
	return
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GenResponse(20000, nil, nil))
}

func GetUserInfo(c *gin.Context) {

	userInfo, e := service.GetUser()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40027, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, userInfo, nil))
	return

}
