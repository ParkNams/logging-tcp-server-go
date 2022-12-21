package apilog

import (
	"time"

	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)



func (apiLogData *ApiLogData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)

	filName := commonConstant.GetLogDirByEnv(commonConstant.GetEnvironment()) + "/api/" + now.Format("2006-01-02")

	if !file.FileExistCheck(filName, commonConstant.FILE_EXTENSION.CSV) {
		file.WriteCSVFile(filName, []string{
			"Date Time",
			"Ip Address",
			"Url",
			"Body",
			"Header",
			"Error",
			"ServerType",
		},0777)
	}
	file.WriteCSVFile(filName,[]string{
		now.Format("2006-01-02T15:04:05"),
		apiLogData.Ip,
		apiLogData.Url,
		apiLogData.Body,
		apiLogData.Header,
		apiLogData.Error,
		apiLogData.ServerType,
	},0777)
}