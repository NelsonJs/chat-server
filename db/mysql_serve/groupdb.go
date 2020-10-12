package mysql_serve

import (
	"chat/config"
	"chat/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Cgroup struct {
	GroupId   string
	Name      string
	Intro     string
	Avatar    string
	OwnerId   string   //群主uid
	Helpers   []string //管理员uid集合
	Members   []string //普通成员
	GroupType int      //群类型 1-公开群 2-私有群 3-聊天室
	Status    int      //1-全体禁言
	Apply     int      //申请入群控制 1-需要验证审核
	Max       int      //群组总共可以容纳的人数
	MaxHelper int      //总共可以容纳的管理员
}

func CreateGroup(ownerId, name, avatar string, groupType int) (bool, error) {
	defer catchErr()
	var builder strings.Builder
	builder.WriteString(ownerId)
	builder.WriteString(name)
	builder.WriteString(strconv.Itoa(groupType))
	builder.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	groupId := utils.Md5(builder.String())
	group := Cgroup{GroupId: groupId, OwnerId: ownerId, Name: name, Avatar: avatar, GroupType: groupType}
	tx := db.Create(&group)
	if tx.Error != nil {
		config.Log.Error(tx.Error.Error())
		return false, tx.Error
	}
	return true, nil
}

func UpdateName(groupId, name string) (bool, error) {
	defer catchErr()
	tx := db.Model(&Cgroup{}).Where("groupid = ?", groupId).Update("name", name)
	if tx.Error != nil {
		config.Log.Error(tx.Error.Error())
		return false, tx.Error
	}
	return true, nil
}

//增加成员
func AddMember(groupId, uid string) (bool, error) {
	defer catchErr()
	var group Cgroup
	tx := db.Where("groupid=?", groupId).First(&group)
	if tx.Error != nil {
		return false, tx.Error
	} else {
		members := group.Members
		if members != nil {
			members = make([]string, 0)
		}
		members = append(members, uid)
		tx := db.Where("groupid = ?", groupId).Update("members", members)
		if tx.Error != nil {
			return false, tx.Error
		}
		return true, nil
	}
}

func catchErr() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
