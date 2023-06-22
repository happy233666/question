package dao

import (
	"fmt"
	"gin/main/massage-board/model"
)

func Insertmassage(m model.Massage) (err error) {
	if m.Mid == "" || m.Uid == "" || m.Lmassage == "" {
		return nil
	}
	_, err = DB.Exec("insert into massageboard(uid,mid,massage)values (?,?,?)", m.Uid, m.Mid, m.Lmassage)
	return err
}
func Searchmid(mid string) (m model.Massage, err error) {
	row := DB.QueryRow("select mid,uid,massage from massageboard where mid = ?", mid)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&m.Uid, &m.Mid, &m.Lmassage)
	return
}
func ChangeMASSAGE(mid string, massage string) {
	sql := "update massageboard set massage =? where mid = ?"
	res, err := DB.Exec(sql, massage, mid)
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
func DeleteMassage(mid string) {
	sql := "DELETE FROM massageboard WHERE mid=?"
	res, err := DB.Exec(sql, mid)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}
	fmt.Println("delete succ:", row)
}
