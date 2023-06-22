package service

import (
	"gin/main/massage-board/dao"
	"gin/main/massage-board/model"
)

func NewComment(m model.Massage) (err error) {
	err = dao.Insertcomment(m)
	return err
}
func ChangeuserCommentByGMID(mid string, comment string) {
	dao.ChangeuserComment(mid, comment)
}
func SearchCommentByMID(mid string) (err error, m model.Massage) {
	m, err = dao.Searchmid(mid)
	return
}
