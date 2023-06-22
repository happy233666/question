package dao

import (
	"fmt"
	"gin/main/massage-board/model"
)

func Insertcomment(m model.Massage) (err error) {
	if m.Mid == "" || m.Gid == "" || m.Comments == "" {
		return nil
	}
	_, err = DB.Exec("insert into comment(gid,mid,comments)values (?,?,?)", m.Gid, m.Mid, m.Comments)
	return err
}
func ChangeuserComment(comments string, mid string) {
	sql := "update comment set comments =? where mid = ? "
	res, err := DB.Exec(sql, comments, mid)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}

	fmt.Println("update succ:", row)

}
