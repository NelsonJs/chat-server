package websocket

import (
	"chat/db/mysql_serve"
	"chat/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

//SendText: 暂时是单聊
func SendText(sendId, receiveId, txt string) (bool, error) {
	client := GetUserClient(receiveId)
	if client == nil {
		return false, errors.New("用户未连接")
	}
	//保存进数据库
	code, err := mysql_serve.SaveRecord(sendId, receiveId, txt, 0)
	if code == -1 {
		return false, err
	}
	//发送
	client.Send <- []byte(txt)
	return true, nil
}

func SendTxt(client *Client, reqModel *models.Req) {
	if client == nil || reqModel == nil {
		return
	}

	c := GetUserClient(strconv.FormatInt(reqModel.ReceiveId, 10))
	if c == nil {
		return
	}
	//保存进数据库
	mysql_serve.SaveRecord(client.UserId, strconv.FormatInt(reqModel.ReceiveId, 10), reqModel.Content, 0)
	resMsg := models.ResMsg{
		SendId:    client.UserId,
		ReceiveId: strconv.FormatInt(reqModel.ReceiveId, 10),
		MsgType:   0,
		Content:   reqModel.Content,
	}
	byt, err := json.Marshal(resMsg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.SendMsg([]byte(byt))
}
