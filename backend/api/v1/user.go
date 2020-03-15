package v1

import (
	"encoding/json"
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

func GetUserAbout(c *gin.Context) {
	about, e := service.GetAbout()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40027, nil, e))
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, about, nil))

}

func EditUserAbout(c *gin.Context) {
	bytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, err))
	}

	u := service.User{}
	if e := json.Unmarshal(bytes, &u); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(20028, nil, err))
		return
	}

	if u.About != "" {
		if e := u.EditAbout(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	} else if u.Password != "" {
		if e := u.ResetPassword(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	} else {
		if e := u.EditUser(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	}

	c.JSON(http.StatusOK, utils.GenResponse(20000, nil, nil))
	return

}
