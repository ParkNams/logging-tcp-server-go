# Local Logging Server

## Log 저장 위치

`$HOME/logs`

## GOLANG 세팅
```
1. go1.18 or go1.19 설치(홈페이지에서 패키지 설치 권장 - brew X) 설치페이지: https://golang.org/dl/
2. go env -w GO111MODULE=auto 설정
3. go mod download 의존성 패키지 설치
4. go get
```

## .env 세팅
```
TCP_PORT=":8000"
ENVIRONMENT="" // local or container
```

## Build 파일 생성
```shell
./script/go-build.sh
# 빌드파일 위치 ./go-build/log-file-go
```

## Docker Container 실행
```shell
# 먼저 .env -> ENVIRONMENT="container" 로 설정
./script/docker-run.sh
```

## 통신 프로토콜
```
TEST_LOG // 테스트용
SYNC_LOG // Batch SyncAll 성능 로그
API_LOG // API Request 로그

PROGRAM_EXIT // 프로그램 종료
```

## 서버 데이터 전달 포맷

```json
{
    "protocol": "" // ... protocol,
    "data": {
        // ... some object
    }
}
```

