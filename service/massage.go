package service

import (
	"gin/main/massage-board/dao"
	"gin/main/massage-board/model"
)

func Creatmassage(m model.Massage) (err error) {
	err = dao.Insertmassage(m)
	return err
}
func SearchmassageByMID(mid string) (err error, m model.Massage) {
	m, err = dao.Searchmid(mid)
	return
}
func ChangeuserMassageByMID(mid string, massage string) {
	dao.ChangeMASSAGE(mid, massage)
}
func DeleteMassageByMID(mid string) {
	dao.DeleteMassage(mid)
}
