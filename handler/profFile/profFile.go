package profFile

import (
	"log"
	"time"

	"github.com/novalagung/gubrak/v2"
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)

type OrderFile struct {
	File []byte
	Order int
	CreatedTime int64
}

var (
	IdleProfFiles = map[string][]OrderFile{}
)


func (profData ProfFileData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)
	fileName := commonConstant.GetLogDirByEnv(commonConstant.GetEnvironment()) + "/prof/" + now.Format("2006-01-02")

	if IdleProfFiles[profData.Uuid] == nil  {
		IdleProfFiles[profData.Uuid] = []OrderFile{
			{
				File: profData.FileByte,
				Order: profData.NowIdx,
				CreatedTime: now.UnixMilli(),
			},
		}
	} else {
		IdleProfFiles[profData.Uuid] = append(IdleProfFiles[profData.Uuid], OrderFile{
			File: profData.FileByte,
			Order: profData.NowIdx,
		})
	}

	 if profData.NowIdx >= profData.MaxIdx {

		makeFile := []byte{}

		orderData := gubrak.From(IdleProfFiles[profData.Uuid]).OrderBy(func (orderb OrderFile) int {
			return orderb.Order
		}).Result().([]OrderFile)

		gubrak.From(orderData).Each(func (data OrderFile) {
			makeFile = append(makeFile, data.File...)
		})
		
		log.Println("write prof")
		file.WriteFile(fileName + string(profData.ProfType),
		makeFile, 0777, commonConstant.FILE_EXTENSION.PROF)
		delete(IdleProfFiles, profData.Uuid)
	}
}