package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var OK = respTemplate{
	Status: 1,
	Info:   "success",
}
var Paramerror = respTemplate{
	Status: 0,
	Info:   "param error",
}
var InternalErr = respTemplate{
	Status: 3,
	Info:   "internal error",
}

func Respok(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}
func RespParamerror(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Paramerror)
}
func RestpInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalErr)
}
func Normalerr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
