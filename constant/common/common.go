package common

import (
	"os"

	"github.com/joho/godotenv"
)

type Protocol struct {
	// logging
	TestLog string
	SyncLog string
	ApiLog string

	ProgramExit string
}

type FileExtension struct {
	TEXT string
	CSV  string
}

var (
	_ = godotenv.Load()
)

const (
	// 서버 연결 방식
	METHOD = "tcp"
)

var (
	// 프로그램 프로토콜
	PROTOCOL = Protocol{
		// logging
		TestLog: "TEST_LOG",
		SyncLog: "SYNC_LOG",
		ApiLog: "API_LOG",
		
		// system
		ProgramExit: "PROGRAM_EXIT",
	}

	// 로깅 관련 프로토콜
	LOGGING_PROTOCOLS = [3]string{
		"TEST_LOG",
		"SYNC_LOG",
		"API_LOG",
	}

	// 시스템 설정 관련 프로토콜
	SYSTEM_PROTOCOLS = [1]string{
		"PROGRAM_EXIT",
	}

	// 파일 확장자
	FILE_EXTENSION = FileExtension{
		CSV:  ".csv",
		TEXT: ".txt",
	}
)

/*
	서버 실행중인 유저 홈 디렉토리 경로
*/
func GetUserHomePath() string {
	home, _ := os.UserHomeDir()
	return home
}

/*
	실행 환경 (local, container)
*/
func GetEnvironment() string {
	var env string = os.Getenv("ENVIRONMENT")
	if env == "" {
		return "local"
	}
	return env
}

/*
	환경변수에 따른 로그 저장 경로
*/
func GetLogDirByEnv(env string) string {
	switch env {
	case "container":
		return "/logging-batch-go/logs"
	case "local":
		return GetUserHomePath() + "/logs"
	default:
		return GetUserHomePath() + "/logs"
	}
}

/*
	tcp 연결 포트
*/
func GetPort() string {
	var port = os.Getenv("TCP_PORT")
	if port == "" {
		return ":8000"
	}
	return port
}
