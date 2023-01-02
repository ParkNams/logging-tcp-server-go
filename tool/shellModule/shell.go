package shellModule

import (
	"os/exec"

	"github.com/joho/godotenv"
	"logFile.com/log-file-go/tool/common"
)

var (
	_ = godotenv.Load()
)

func OpenCpuProfHttp() error {
	cmd := exec.Command("sh", "/logging-batch-go/script/open-cpu-prof.sh")

	runErr := cmd.Run()

	if common.ErrorLogging(runErr) {
		return runErr
	}
	return nil
}

func OpenMemProfHttp() error {
	cmd := exec.Command("sh", "/logging-batch-go/script/open-mem-prof.sh")

	runErr := cmd.Run()

	if common.ErrorLogging(runErr) {
		return runErr
	}
	return nil
}
