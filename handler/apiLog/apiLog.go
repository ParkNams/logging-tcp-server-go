package apilog

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/awsModule"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)

var (
	_ = godotenv.Load()
)

func (apiLogData *ApiLogData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)
	fileOriginName := apiLogData.ServerType + "_" + apiLogData.User + "_" +
		now.Format("2006-01-02")
	filName := commonConstant.GetLogDirByEnv(commonConstant.GetEnvironment()) +
		"/api/" +
		fileOriginName

	if !file.FileExistCheck(filName, commonConstant.FILE_EXTENSION.CSV) {
		file.WriteCSVFile(filName, []string{
			"Date Time",
			"Ip Address",
			"Url",
			"Body",
			"Header",
			"Error",
			"ServerType",
		}, 0777)
	}
	file.WriteCSVFile(filName, []string{
		now.Format("2006-01-02T15:04:05"),
		apiLogData.Ip,
		apiLogData.Url,
		apiLogData.Body,
		apiLogData.Header,
		apiLogData.Error,
		apiLogData.ServerType,
	}, 0777)

	fileBody := file.GetFile(filName, commonConstant.FILE_EXTENSION.CSV)

	if fileBody != nil {
		log.Println("s3 upload api logging")
		awsModule.UploadS3(
			os.Getenv("LOGGING_BUCKET"),
			"logging/api/"+
				fileOriginName+
				commonConstant.FILE_EXTENSION.CSV,
			string(fileBody),
			"multipart/formed-data",
		)
		// file.RemoveFile(filName, commonConstant.FILE_EXTENSION.CSV)
	}
}
