package utils

import (
	"fmt"
	"time"
)

const (
	HandleErrorLevelInfo = iota
	HandleErrorLevelWarn
	HandleErrorLevelError
)
const (
	HandleDescErrorTypeShow    = iota //仅打印错误信息
	HandleDescErrorTypeLog            //打印错误信息并写入日志
	HandleDescErrorTypeSendMsg        //打印错误信息发生邮件通知
)

// HandleError 全局处理错误
func HandleError(err error, level int, descErrorType int, args ...interface{}) {
	if err == nil {
		return
	}
	//仅打印数据
	fmt.Printf("%s\t%v%v\n", time.Now().Format("2006-01-02 15:04:05"), err, args)

	//按错误类别处理...

}
