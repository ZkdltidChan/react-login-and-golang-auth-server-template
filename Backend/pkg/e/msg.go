package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "Invaild params",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token error",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token timeout",
	ERROR_AUTH_TOKEN:               "Token create error",
	ERROR_AUTH:                     "Token error 2",
	ERROR_USER_NOT_FOUND:           "ERROR_USER_NOT_FOUND",
	ERROR_HASH_PASSWORD:            "ERROR_HASH_PASSWORD",
	ERROR_UNDIFINE: 	"ERROR_UNDIFINE",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
