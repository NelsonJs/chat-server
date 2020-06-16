package mysql_serve

import (
	"chat/models"
	"database/sql"
	"encoding/json"
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

func activeList() (error, []*models.ResDynamic) {
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

func NearDynamic() (error, []*models.ResDynamic) {
	stmt, err := db.Prepare("select * from dynamic order by time desc")
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
		err = rows.Scan(&bean.Id, &bean.Uid, &bean.Title, &bean.Img, &bean.Like, &bean.Comment_id, &bean.Loc, &bean.Lat, &bean.Lng, &bean.Time, &bean.Res_img)
		if err != nil {
			return err, nil
		}
		list = append(list, &bean)
	}
	return nil, list
}

func PublishDynamic(uid string,title string,imgIds []int64) (int64,error) {
	tx,err := db.Begin()
	if err != nil {
		return 0, err
	}
	var queryStr string

	queryCon := make([]int64,0)
	for i:=0; i< len(imgIds);i++ {
		if i < len(imgIds)-1 {
			queryStr += "id=? or "
		} else {
			queryStr += "id=?"
		}
		queryCon = append(queryCon,imgIds[i])
	}
	fmt.Printf("id个数: %d 查询语句：%s 查询参数：%s ",len(imgIds),queryStr,queryCon)
	stmt,err := tx.Prepare("select path from imgs where "+queryStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var rows *sql.Rows
	if len(queryCon) == 1 {
		rows,err = stmt.Query(queryCon[0])
	} else if len(queryCon) == 2 {
		rows,err = stmt.Query(queryCon[0],queryCon[1])
	} else if len(queryCon) == 3 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2])
	} else if len(queryCon) == 4 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2],queryCon[3])
	} else if len(queryCon) == 5 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2],queryCon[3],queryCon[4])
	} else if len(queryCon) == 6 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2],queryCon[3],queryCon[4],queryCon[5])
	} else if len(queryCon) == 7 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2],queryCon[3],queryCon[4],queryCon[5],queryCon[6])
	} else if len(queryCon) == 8 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2],queryCon[3],queryCon[4],queryCon[5],queryCon[6],queryCon[7])
	} else if len(queryCon) == 9 {
		rows,err = stmt.Query(queryCon[0],queryCon[1],queryCon[2],queryCon[3],queryCon[4],queryCon[5],queryCon[6],queryCon[7],queryCon[8])
	}

	if err != nil {
		fmt.Printf("查询语句错误：%s \n",err.Error())
		return 0, err
	}
	var path string
	paths := make([]string,0)
	for rows.Next() {
		err = rows.Scan(&path)
		if err != nil {
			fmt.Println("查询出错：",err.Error())
		} else {
			fmt.Println("图片地址--》",path)
			paths = append(paths,path)
		}
	}
	stmt,err = db.Prepare("insert into dynamic(uid,title,time,res_img)values(?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	byt,err := json.Marshal(paths)
	if err != nil {
		return 0,err
	}
	result,err := stmt.Exec(uid,title,time.Now().Unix(),byt)
	if err != nil  {
		return 0, err
	}
	return result.LastInsertId()
}

func UploadAvatar(path string) (int64, error) {
	stmt, err := db.Prepare("insert into imgs(path) value(?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result,err := stmt.Exec(path)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func AddImg(path []string) ([]int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare("insert into imgs(path) value(?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res = make([]int64, 0)
	for i := 0; i < len(path); i++ {
		result, err := stmt.Exec(path[i])
		if err != nil {
			tx_rollback(err, tx)
		} else {
			id, err := result.LastInsertId()
			if err != nil {
				fmt.Println("插入出错：", err.Error())
			} else {
				res = append(res, id)
			}

		}
	}
	err = tx.Commit()
	if err != nil {
		tx_rollback(err, tx)
		return res[0:0], err
	}
	return res, nil
}


func tx_rollback(err error, tx *sql.Tx) {
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("事务回滚失败")
			return
		}
	}
}
