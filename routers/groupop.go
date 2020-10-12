package routers

import (
	"chat/db/mysql_serve"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(c *gin.Context) {
	var group mysql_serve.Cgroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if group.OwnerId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ownerId缺失",
		})
		return
	}
	if group.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "未设置群组名称",
		})
		return
	}
	_, err := mysql_serve.CreateGroup(group.OwnerId, group.Name, group.Avatar, group.GroupType)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "创建成功",
		})
	}
}

func UpdateName(c *gin.Context) {
	var g mysql_serve.Cgroup
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	_, err := mysql_serve.UpdateName(g.GroupId, g.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "更新成功",
		})
	}
}

//增加管理员
func AddManager(c *gin.Context) {

}

//主动增加成员
func AddMember(c *gin.Context) {

}

//移除群组中人员
func RemoveMember(c *gin.Context) {

}

//添加群组头像
func AddAvatar(c *gin.Context) {

}

//解散群组
func Del(c *gin.Context) {

}

//转移群组
func Transfer(c *gin.Context) {

}

//申请加入群组
func Join(c *gin.Context) {

}

//退出群组
func Exit(c *gin.Context) {

}
