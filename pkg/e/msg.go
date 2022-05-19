package e

var MsgFlags = map[int]string{
	SUCCESS:                      "ok",
	ERROR:                        "fail",
	INVALID_PARAMS:               "请求参数错误",
	RECORD_NOT_FOUND:             "数据库找不到记录",
	TOKEN_NOT_FOUND:              "找不到token",
	INVALID_TOKEN_PREFIX:         "无效的token前缀",
	INVALID_TOKEN:                "无效的token",
	TOKEN_IS_EXPIRED:             "token已失效",
	NO_AVALIABLE_USER_FROM_TOKEN: "找不到对应的用户信息",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
