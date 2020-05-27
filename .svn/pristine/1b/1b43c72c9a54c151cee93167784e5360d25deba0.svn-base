package mysql_serve

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMySQL() {
	var err error
	db, err = sql.Open("mysql", "root:6678510jk@tcp(127.0.0.1:3306)/demo?charset=utf8mb4")
	if err != nil {
		fmt.Println("打开数据库连接失败", err)
		return
	}
}

//GetConversations: 获取uid拥有的所有会话
func GetConversations(uid string) {
	stmt, err := db.Prepare("select receive_id as cids from msg where send_id=? union select send_id from msg where receive_id=?")
	if err != nil {
		fmt.Println("获取会话列表错误", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid, uid)
	defer rows.Close()
	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("查找出现异常", err)
			return
		}
		fmt.Println("会话id：", id)
	}
}

func Register(userName, pwd string) (code int64, msg string) {
	stmt, err := db.Prepare("insert into user(nick_name,pwd)values(?,?)")
	if err != nil {
		return -1, err.Error()
	}
	defer stmt.Close()
	result, err := stmt.Exec(userName, pwd)
	if err != nil {
		return -1, err.Error()
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err.Error()
	}
	if affected > 0 {
		id, err := result.LastInsertId()
		if err != nil {
			return -1, err.Error()
		} else {
			fmt.Println("最后影响的id为:", id)
			return id, "插入成功"
		}

	}
	return -1, "插入失败"
}
