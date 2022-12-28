package profFile

import (
	"log"
	"time"

	"github.com/novalagung/gubrak/v2"
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/constant/localCache"
	"logFile.com/log-file-go/structures"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)

// var (
// 	IdleProfFiles = map[string][]OrderFile{}
// )

func (profData ProfFileData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)
	fileName := commonConstant.GetLogDirByEnv(commonConstant.GetEnvironment()) + "/prof/" + now.Format("2006-01-02")

	if localCache.IdleProfFiles[profData.Uuid] == nil  {
		localCache.IdleProfFiles[profData.Uuid] = []structures.OrderFile{
			{
				File: profData.FileByte,
				Order: profData.NowIdx,
				CreatedTime: now.UnixMilli(),
			},
		}
	} else {
		localCache.IdleProfFiles[profData.Uuid] = append(localCache.IdleProfFiles[profData.Uuid], structures.OrderFile{
			File: profData.FileByte,
			Order: profData.NowIdx,
		})
	}

	 if profData.NowIdx >= profData.MaxIdx {

		makeFile := []byte{}

		orderData := gubrak.From(localCache.IdleProfFiles[profData.Uuid]).OrderBy(func (orderb structures.OrderFile) int {
			return orderb.Order
		}).Result().([]structures.OrderFile)

		gubrak.From(orderData).Each(func (data structures.OrderFile) {
			makeFile = append(makeFile, data.File...)
		})
		
		log.Println("write prof")
		file.WriteFile(fileName + string(profData.ProfType),
		makeFile, 0777, commonConstant.FILE_EXTENSION.PROF)
		delete(localCache.IdleProfFiles, profData.Uuid)
	}
}
