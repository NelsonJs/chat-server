package mysql_serve

import (
	"chat/models"
	"database/sql"
	"fmt"
	"time"

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

func SaveRecord(sendId, receiveId, content string, msgType int) (int, error) {
	stmt, err := db.Prepare("insert into msg(send_id,receive_id,msg_type,content,create_time)values(?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(sendId, receiveId, msgType, content, time.Now().Unix())
	if err != nil {
		return -1, err
	}
	affectedNum, err := result.RowsAffected()
	if affectedNum > 0 {
		return 1, nil
	}
	return -1, err
}

//GetConversations: 获取uid拥有的所有会话
func GetConversations(uid string) []*models.UserInfo {
	ids := make([]*models.UserInfo, 0)
	stmt, err := db.Prepare("select receive_id as cids from msg where send_id=? union select send_id from msg where receive_id=?")
	if err != nil {
		fmt.Println("获取会话列表错误", err)
		return ids
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid, uid)
	defer rows.Close()
	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("查找出现异常", err)
			return ids
		}
		fmt.Println("会话id：", id)
		code, b := GetUserInfo(id)
		if code == 1 {
			ids = append(ids, b)
		}

	}
	return ids
}

func GetUserInfo(uid string) (code int, info *models.UserInfo) {
	stmt, err := db.Prepare("select uid,nick_name,gender from user where uid=?")
	if err != nil {
		fmt.Println(err.Error())
		return -1, nil
	}
	defer stmt.Close()
	row := stmt.QueryRow(uid)
	userInfo := models.UserInfo{}
	err = row.Scan(&userInfo.Uid, &userInfo.UserName, &userInfo.Gender)
	if err != nil {
		fmt.Println("查找出现异常", err)
		return -1, nil
	}
	return 1, &userInfo
	// rows, err := stmt.Query(uid)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return -1, nil
	// }
	// defer rows.Close()
	// userInfo := models.UserInfo{}
	// for rows.Next() {
	// 	err = rows.Scan(&userInfo.Uid, &userInfo.UserName, &userInfo.Gender)
	// 	if err != nil {
	// 		fmt.Println("查找出现异常", err)
	// 		return -1, nil
	// 	}
	// 	return 1, &userInfo
	// }
	// return -1, nil
}

func GetChatRecord(selfId, otherId int64) (int, string, []*models.Message) {
	stmt, err := db.Prepare("select send_id,receive_id,content from msg where send_id=? and receive_id=? or send_id=? and receive_id=? order by create_time desc")
	if err != nil {
		return -1, err.Error(), nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(selfId, otherId, otherId, selfId)
	if err != nil {
		return -1, err.Error(), nil
	}
	list_ := make([]*models.Message, 0)
	for rows.Next() {
		data := models.Message{}
		err = rows.Scan(&data.SendId, &data.ReceiveId, &data.Content)
		if err != nil {
			return -1, err.Error(), nil
		}
		fmt.Println(data.Content)
		list_ = append(list_, &data)
	}
	return 1, "", list_
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

func Login(userName, pwd string) (code int64, msg string) {
	if userName == "" || pwd == "" {
		return -1, "用户名或密码不能为空"
	}
	stmt, err := db.Prepare("select uid from user where nick_name=? and pwd=?")
	if err != nil {
		return -1, err.Error()
	}
	defer stmt.Close()
	row := stmt.QueryRow(userName, pwd)
	var uid int64
	err = row.Scan(&uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return -2, "该用户还未注册"
		}
		return -1, err.Error()
	}
	fmt.Println("登录成功")
	return uid, "登录成功"
}

func NearDynamic() (error, []*models.ResDynamic) {
	stmt, err := db.Prepare("select * from active")
	if err != nil {
		return err, nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return err, nil
	}
	list := make([]*models.ResDynamic, 0)
	for rows.Next() {
		bean := models.ResDynamic{}
		err = rows.Scan(&bean.Id, &bean.Uid, &bean.Title, &bean.Description, &bean.Img, &bean.Gender, &bean.Begin, &bean.Loc, &bean.Lng, &bean.Lat, &bean.People_num, &bean.People_total_num, &bean.Like, &bean.Comment_num, &bean.Comment_id)
		if err != nil {
			return err, nil
		}
		list = append(list, &bean)
	}
	return nil, list
}
