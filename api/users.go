package api

import (
	"database/sql"
	"gin/main/massage-board/model"
	"gin/main/massage-board/service"
	"gin/main/massage-board/util"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) {
	id := c.PostForm("id")
	userName := c.PostForm("username")
	password := c.PostForm("password")
	err, u := service.SearchuserByID(id)
	if err != nil && err != sql.ErrNoRows {
		util.RestpInternalErr(c)
		return
	}
	if u.Id != "" {
		util.Normalerr(c, 3, "账号已存在")
		return
	}
	err = service.Creatuser(model.User{
		Id:       id,
		Username: userName,
		Password: password,
	})
	if id == "" || userName == "" || password == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		util.RestpInternalErr(c)
		return
	}
	util.Respok(c)

}
func Login(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	err, u := service.SearchuserByID(id)
	if id == "" || password == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "用户不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if u.Password != password {
		util.Normalerr(c, 4, "密码错误")
		return
	}
	c.SetCookie(id, password, 0, "", "", false, false)
	util.Respok(c)
}
func ChangePassword(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	err, u := service.SearchuserByID(id)
	if id == "" || password == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "用户不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if u.Password == password {
		util.Normalerr(c, 5, "密码重复")
		return
	}
	service.ChangeuserPasswordByID(id, password)
	util.Respok(c)
}
