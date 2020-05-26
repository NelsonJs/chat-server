package websocket

import (
	"chat/models"
	"encoding/json"
)

func SendTxt(client *Client, reqModel *models.Req) {
	if client == nil || reqModel == nil {
		return
	}
	isRegister := UserIsRegister(client, reqModel)
	if !isRegister { //如果没有注册
		res := &models.Res{Code: -1, Msg: "该用户不存在"}
		resData, _ := json.Marshal(res)
		client.SendMsg(resData)
		return
	}
	//获取用户类型（单聊，群聊）
	//需要获取对应的用户链接
	//保存消息到消息表
	//在线->发送
	//不在线->保存消息到离线表->发送离线推送
	c := GetUserClient(reqModel.OtherUserId)
	if c != nil {
		client.SendMsg(reqModel.MsgByte)
		c.SendMsg(reqModel.MsgByte)
	} else {
		res := &models.Res{Code: -1, Msg: "没有找到" + reqModel.OtherUserId}
		resData, _ := json.Marshal(res)
		client.SendMsg(resData)
	}

	// v, b := reqModel.Msg.([]byte)
	// if b {
	// 	//需要获取对应的用户链接
	// 	c := GetUserClient(reqModel.OtherUserId)
	// 	if c != nil {
	// 		c.SendMsg(v)
	// 	} else {
	// 		res := &models.Res{Code: -1, Msg: "内部数据异常"}
	// 		resData, _ := json.Marshal(res)
	// 		client.SendMsg(resData)
	// 	}
	// } else {
	// 	res := &models.Res{Code: -1, Msg: "数据转换异常"}
	// 	resData, _ := json.Marshal(res)
	// 	client.SendMsg(resData)
	// }
}
