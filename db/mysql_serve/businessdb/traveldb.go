package businessdb

import (
	"chat/db/mysql_serve"
	"chat/utils"
	"encoding/json"
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


