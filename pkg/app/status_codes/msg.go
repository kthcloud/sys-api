package status_codes

var MsgFlags = map[int]string{
	Unknown:       "unknown",
	Success:       "success",
	Error:         "error",
	InvalidParams: "invalidParams",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
