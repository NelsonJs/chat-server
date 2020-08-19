package mysql_serve

import (
	"chat/logger"
	"chat/models"
)

func GetUsers() []*models.UserInfo_{
	list := make([]*models.UserInfo_,0)
	stmt,err := db.Prepare("select * from user order by login_time desc limit 10")
	if err != nil {
		logger.LogManager.Error(err.Error())
		return list
	}
	defer stmt.Close()
	rows,err := stmt.Query()
	if err != nil {
		logger.LogManager.Error(err.Error())
		return list
	}
	for rows.Next() {
		bean := models.UserInfo_{}
		err = rows.Scan(&bean.Uid,&bean.Nick_name,&bean.Pwd,&bean.Phone,&bean.Gender,&bean.Year_old,&bean.Avatar,&bean.Create_time,&bean.Login_time,&bean.Logout_time,&bean.Status)
		if err != nil {
			logger.LogManager.Error(err.Error())
			return list
		}
		list = append(list,&bean)
	}
	return list
}
