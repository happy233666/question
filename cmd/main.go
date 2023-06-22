package main

import (
	"gin/main/massage-board/api"
	"gin/main/massage-board/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
