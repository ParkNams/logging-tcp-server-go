package shellModule

import (
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/common"
)

var (
	_ = godotenv.Load()
)

func OpenProfHttp(openType string, port string) error {
	filePath := commonConstant.GetLogDirByEnv(os.Getenv("ENVIRONMENT")) + "/" +
		openType

	cmd := exec.Command("sh", "./script/open-prof.sh", filePath, "openHttp.prof", port)

	runErr := cmd.Run()

	if common.ErrorLogging(runErr) {
		return runErr
	}
	return nil
}
