package group

type Group struct {
	GroupId   string
	Name      string
	Intro     string
	Avatar    string
	OwnerId   string   //群主用户uid
	Helpers   []string //管理员uid集合
	GroupType int      //群类型 1-公开群 2-私有群 3-聊天室
	Status    int      //1-全体禁言
	Apply     int      //申请入群控制 1-需要验证审核
	Max       int      //群组总共可以容纳的人数
	MaxHelper int      //总共可以容纳的管理员
}

// 创建群组
func CreateGroup(ownerId, name, avatar string, groupType int) {

}

// 更新群组信息
func (group *Group) UpdateGroupInfo() {

}

// 增加管理员或成员
func (group *Group) AddMember(members, helpers []string) {

}

// 移除群组中人员
func (group *Group) RemoveMember(uid, reason string) {

}

// 申请加入群组
func (group *Group) Join() {

}

// 退出群组
func (group *Group) Exit() {

}
