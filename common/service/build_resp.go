package service

import (
	"log"
)

//BuildResponse 根据err返回对应的Response
func BuildResponse(err error) (statusCode int32, statusMsg string) {
	if err != nil {
		//这里暂时还没协商出errno code，所以错误先默认为
		log.Println(err)
		statusCode = -1
		statusMsg = "fail"
	} else {
		statusCode = 0
		statusMsg = "success"
	}
	return
}
