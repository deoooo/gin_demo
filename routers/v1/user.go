package v1

import (
	"github.com/deoooo/gin_demo/models"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	code := e.SUCCESS
	data := make(map[string]interface{})

	ok, u := models.CheckPassword(username, password)
	if ok {
		token, _ := util.GenerateToken(u.ID)
		data["token"] = token
	} else {
		code = e.AuthFail
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Hello(c *gin.Context) {
	code := e.SUCCESS
	userId, exit := c.Get("userId")
	data := make(map[string]interface{})
	if exit {
		data["userId"] = userId
	} else {
		code = e.AuthFail
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
