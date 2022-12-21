package main

import (
	"log"
	"runtime"

	"logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/controller"
	"logFile.com/log-file-go/tool/tcp"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	l := tcp.GetListener(common.METHOD, common.GetPort())

	defer l.Close()
	defer log.Println("process exit")

	for {
		log.Println("exec connect")
		if controller.Controller(l) {
			break
		}
		continue
	}
}
