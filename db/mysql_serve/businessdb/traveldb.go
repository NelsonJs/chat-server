package businessdb

import (
	"chat/constants"
	"chat/db/mysql_serve"
	"chat/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Travel struct {
	Id int64 `json:"-"`
	Tid string `json:"tid"`
	Ttype string `json:"ttype"` //出行 活动
	Car string `json:"car"`
	Cartype string `json:"cartype"`
	Carnum int `json:"carnum"'`
	Uid string `json:"uid"`
	Title string `json:"title"`
	Starttime int64 `json:"starttime"`
	Startloc string `json:"startloc"`
	Driveloc string `json:"driveloc"`
	Endloc string `json:"endloc"`
	Loclat float64 `json:"loclat"`
	Loclng float64 `json:"loclng"`
	Price string `json:"price"`
	Total int `json:"total"`
	Curnum int `json:"curnum"`
	Description string `json:"description"`
	Members json.RawMessage `json:"members"`
	Status int `json:"status"`
	Createtime int64 `json:"createtime"`
}

type TravelOut struct {
	StartPlace string `json:"start_place"`
	EndPlace string `json:"end_place"`
	TravelType string `json:"travel_type"`
	Travels []*Travel `json:"travels"`
}

type NameUser struct {
	Uid string `json:"uid"`
	Avatar string `json:"avatar"`
}

func PublishTravel(travel *Travel) error {
	tid := utils.Md5WithTime(travel.Title)
	travel.Tid = tid
	travel.Createtime = time.Now().Unix()
	tx := mysql_serve.Db.Create(&travel)
	return tx.Error
}

func GetTravelList() (error,[]*TravelOut) {
	outTravels := make([]*TravelOut,0)
	var travels []*Travel
	tx := mysql_serve.Db.Order("createtime desc").Find(&travels)
	if travels != nil && len(travels) > 0 {
		sameTravel := make(map[string][]*Travel,0)
		for _,v := range travels {
			var buf strings.Builder
			buf.WriteString(v.Ttype)
			buf.WriteString(v.Startloc)
			buf.WriteString(v.Endloc)
			if _, ok := sameTravel[buf.String()];ok {
				sameTravel[buf.String()] = append(sameTravel[buf.String()],v)
			} else {
				tempTravels := make([]*Travel,0)
				tempTravels = append(tempTravels,v)
				sameTravel[buf.String()] = tempTravels
			}
		}
		for _,v := range sameTravel {
			var t TravelOut
			t.StartPlace = v[0].Startloc
			t.EndPlace =v[0].Endloc
			t.TravelType = v[0].Ttype
			t.Travels = v
			outTravels = append(outTravels,&t)
		}
	}
	return tx.Error,outTravels
}

func JoinTravel(tid, uid string) error {
	var t Travel
	tx := mysql_serve.Db.Where("tid = ?",tid).First(&t)
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println(string(t.Members))
	var nameUsers []NameUser
	err := json.Unmarshal(t.Members,&nameUsers)
	if err != nil {
		return err
	}
	for _,v := range nameUsers {
		if v.Uid == uid {
			return constants.ErrUserInTravelExists
		}
	}
	if t.Total <= len(nameUsers) {
		return constants.ErrTravelValueIsOut
	}
	var u mysql_serve.User
	tx = mysql_serve.Db.Where("uid = ?",uid).First(&u)
	if tx.Error == nil {
		var nu NameUser
		nu.Avatar = u.Avatar
		nu.Uid = uid
		nameUsers = append(nameUsers, nu)
		byt,err := json.Marshal(&nameUsers)
		if err == nil {
			tx = mysql_serve.Db.Model(t).Where("tid = ?",tid).Update("members",byt)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}
	return nil
}

func ExitTravel(tid, uid string) error {
	var t Travel
	tx := mysql_serve.Db.Where("tid = ?",tid).First(&t)
	if tx.Error != nil {
		return tx.Error
	}
	var nameUsers []NameUser
	err := json.Unmarshal(t.Members,&nameUsers)
	if err != nil {
		return err
	}
	var joined bool
	for k,v := range nameUsers {
		if v.Uid == uid {
			joined = true
			//设置变量后，同时清除参加信息
			nameUsers = append(nameUsers[:k], nameUsers[k+1:]...)
			break
		}
	}
	if !joined {
		return constants.ErrNotJoinTravel
	}
	byt,err := json.Marshal(&nameUsers)
	if err == nil {
		tx = mysql_serve.Db.Model(t).Where("tid = ?",tid).Update("members",byt)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}


