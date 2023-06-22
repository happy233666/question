package dao

import (
	"fmt"
	"gin/main/massage-board/model"
)

func Insertuser(u model.User) (err error) {
	if u.Username == "" || u.Password == "" {
		return nil
	}
	_, err = DB.Exec("insert into user(id,username,password)values (?,?,?)", u.Id, u.Username, u.Password)
	return err
}
func Searchid(id string) (u model.User, err error) {
	row := DB.QueryRow("select id,username,password from user where id = ?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.Username, &u.Password)
	return
}
func ChangeDB(id string, password string) {
	sql := "update user set password =? where id = ?"
	res, err := DB.Exec(sql, password, id)
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
