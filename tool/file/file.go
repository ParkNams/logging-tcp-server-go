package file

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"

	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/common"
)

/*
	파일 존재 여부 체크
*/
func FileExistCheck(fileName string, fileType string) bool {
	rdByte, err := os.ReadFile(fileName + fileType)
	log.Printf("byte: %d, err: %v", len(rdByte), err)
	return len(rdByte) > 0 && err == nil
}

/*
	csv 파일 쓰기
*/
func WriteCSVFile(fileName string, logText []string, fileAuth int) {
	fileNameSpit := strings.Split(fileName, "/")
	filePath := strings.Join(fileNameSpit[0:len(fileNameSpit)-1], "/")
	log.Println(filePath)
	err := os.MkdirAll(filePath, os.FileMode(fileAuth))

	common.CheckErr(err)

	file, err := os.OpenFile(
		fileName+commonConstant.FILE_EXTENSION.CSV,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(fileAuth))

	common.CheckErr(err)

	defer file.Close()

	wr := csv.NewWriter(bufio.NewWriter(file))
	wr.Write(logText)
	wr.Flush()
}

/*
	text 파일 쓰기
*/
func WriteTextFile(fileName string, fileByte []byte, fileAuth int) {
	fileNameSpit := strings.Split(fileName, "/")
	filePath := strings.Join(fileNameSpit[0:len(fileNameSpit)-1], "/")
	log.Println(filePath)
	err := os.MkdirAll(filePath, os.FileMode(fileAuth))

	common.CheckErr(err)

	file, err := os.OpenFile(
		fileName+commonConstant.FILE_EXTENSION.TEXT,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(fileAuth))

	common.CheckErr(err)

	defer file.Close()

	file.Write(fileByte)
}
