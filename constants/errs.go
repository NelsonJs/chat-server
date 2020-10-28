package constants

import "errors"

var OK = 1

var ErrParameters = 10000
var ErrParseData = 10001

var ErrNotRegister = 20000
var ErrHasRegistered = 20001
var ErrNotReciver = 20002

var ErrArgumentNotExists = errors.New("参数缺失")
var ErrUserNotExists = errors.New("用户名或密码错误")
var ErrUserHasRegister = errors.New("该账号已被注册")
