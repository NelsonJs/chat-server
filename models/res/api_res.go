package res

import "chat/models"

type Conversation struct {
	Data []models.UserInfo
}

type ChatRecord struct {
	Code int
	Msg string
	Data []interface{}
}

type Register struct {
	Uid int
	Username string
	Pwd string
}

type Fail struct {
	Code int
	Msg string
}

