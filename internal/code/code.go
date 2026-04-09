package code

const (
	CodeSuccess        = 0
	CodeInvalidParams  = 4001
	CodeUnauthorized   = 4002
	CodeNotFound       = 4004
	CodeInternalServer = 5000
	CodeDBError        = 5001
)

var codeMessage = map[int]string{
	CodeSuccess:        "success",
	CodeInvalidParams:  "invalid params",
	CodeUnauthorized:   "unauthorized",
	CodeNotFound:       "resource not found",
	CodeInternalServer: "internal server error",
	CodeDBError:        "database error",
}

func GetMessage(code int) string {
	msg, ok := codeMessage[code]
	if !ok {
		return "unknown error"
	}
	return msg
}
