package utils

import (
	"encoding/json"
	"syscall"
)

/****************************************************************
** Serialization
********/

func GetJsonString(v interface{}) string {
	result, _ := json.MarshalIndent(v, "", "  ")
	return string(result)
}

/****************************************************************
** Path utils
********/

/****************************************************************
** Time utils
********/

func NowMicroseconds() int64 {
	tv := syscall.Timeval{}
	syscall.Gettimeofday(&tv)
	return (int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3)
}
