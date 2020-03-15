package v1

import (
	"github.com/coding-codes/service"
	"github.com/coding-codes/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCategory(c *gin.Context) {
	categories, e := service.Category{}.GetAll()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(-1, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, categories, nil))
}

func CreateCategory(c *gin.Context) {
	category := &service.Category{}
	if e := c.ShouldBindJSON(category); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(-1, nil, e))
		return
	}

	cc, e := category.Create()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(-1, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, cc, nil))
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := &service.Category{ID: id}

	ca, e := r.GetOne()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(-1, nil, e))
		return
	}

	if e := ca.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(-1, nil, e))
		return
	}

	c.JSON(http.StatusOK, utils.GenResponse(20000, ca, nil))
}

func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ca := service.Category{ID: id}

	if e := c.ShouldBindJSON(&ca); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(-1, nil, e))
		return
	}

	if e := ca.Edit(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(-1, nil, e))
		return
	}

	c.JSON(http.StatusOK, utils.GenResponse(20000, ca, nil))

}
