package util

import (
	"log"
	"time"

	"logFile.com/log-file-go/handler/profFile"
	"logFile.com/log-file-go/tool/common"
)

func RemoveTrashData() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)
	now := time.Now().In(loc).UnixMilli()

	if len(profFile.IdleProfFiles) == 0 {
		log.Println("profile data empty")
	}else {
		for key,data := range profFile.IdleProfFiles {
			if len(data) > 0 && now - data[0].CreatedTime > 6000 {
				delete(profFile.IdleProfFiles,key)
				log.Println("trash data removed")
			}
		}
	}
}